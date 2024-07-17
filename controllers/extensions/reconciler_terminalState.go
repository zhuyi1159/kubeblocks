/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package extensions

import (
	"fmt"

	extensionsv1alpha1 "github.com/apecloud/kubeblocks/apis/extensions/v1alpha1"
	intctrlutil "github.com/apecloud/kubeblocks/pkg/controllerutil"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/apecloud/kubeblocks/pkg/controller/kubebuilderx"
	"github.com/apecloud/kubeblocks/pkg/controller/model"
	"k8s.io/apimachinery/pkg/api/meta"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type terminalStateReconciler struct {
	stageCtx
}

func (r *terminalStateReconciler) PreCondition(tree *kubebuilderx.ObjectTree) *kubebuilderx.CheckResult {
	if tree.GetRoot() == nil || model.IsObjectDeleting(tree.GetRoot()) {
		return kubebuilderx.ResultUnsatisfied
	}

	return kubebuilderx.ResultSatisfied
}

func (r *terminalStateReconciler) Reconcile(tree *kubebuilderx.ObjectTree) (*kubebuilderx.ObjectTree, error) {
	r.process(func(addon *extensionsv1alpha1.Addon) {
		r.reqCtx.Log.V(1).Info("terminalStateStage", "phase", addon.Status.Phase)
		patchPhaseNCondition := func(phase extensionsv1alpha1.AddonPhase, reason string) {
			r.reqCtx.Log.V(1).Info("patching status", "phase", phase)
			patch := client.MergeFrom(addon.DeepCopy())
			addon.Status.Phase = phase
			addon.Status.ObservedGeneration = addon.Generation

			meta.SetStatusCondition(&addon.Status.Conditions, metav1.Condition{
				Type:               extensionsv1alpha1.ConditionTypeSucceed,
				Status:             metav1.ConditionTrue,
				ObservedGeneration: addon.Generation,
				Reason:             reason,
				LastTransitionTime: metav1.Now(),
			})

			if err := r.reconciler.Status().Patch(r.reqCtx.Ctx, addon, patch); err != nil {
				r.setRequeueWithErr(err, "")
				return
			}
			r.reconciler.Event(addon, corev1.EventTypeNormal, reason,
				fmt.Sprintf("Progress to %s phase", phase))
			r.setReconciled()
		}

		// transit to enabled or disable phase
		switch addon.Status.Phase {
		case "", extensionsv1alpha1.AddonDisabling:
			patchPhaseNCondition(extensionsv1alpha1.AddonDisabled, AddonDisabled)
			return
		case extensionsv1alpha1.AddonEnabling:
			patchPhaseNCondition(extensionsv1alpha1.AddonEnabled, AddonEnabled)
			return
		}
	})

	return tree, nil
}

func NewTerminalStateReconciler(reqCtx intctrlutil.RequestCtx, buildStageCtx func() stageCtx) kubebuilderx.Reconciler {
	return &terminalStateReconciler{
		stageCtx: buildStageCtx(),
	}
}

var _ kubebuilderx.Reconciler = &terminalStateReconciler{}
