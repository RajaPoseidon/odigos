package gateway

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	odigosv1 "github.com/odigos-io/odigos/api/odigos/v1alpha1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func syncService(gateway *odigosv1.CollectorsGroup, ctx context.Context, c client.Client, scheme *runtime.Scheme) (*v1.Service, error) {
	logger := log.FromContext(ctx)
	gatewaySvc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      gateway.Name,
			Namespace: gateway.Namespace,
			Labels:    commonLabels,
		},
	}

	if err := ctrl.SetControllerReference(gateway, gatewaySvc, scheme); err != nil {
		logger.Error(err, "failed to set controller reference")
		return nil, err
	}

	result, err := controllerutil.CreateOrPatch(ctx, c, gatewaySvc, func() error {
		updateGatewaySvc(gatewaySvc)
		return nil
	})

	if err != nil {
		logger.Error(err, "failed to create or patch gateway service")
		return nil, err
	}

	logger.V(0).Info("gateway service synced", "result", result)
	return gatewaySvc, nil
}

func updateGatewaySvc(svc *v1.Service) {
	svc.Spec.Ports = []v1.ServicePort{
		{
			Name:       "otlp",
			Protocol:   "TCP",
			Port:       4317,
			TargetPort: intstr.FromInt(4317),
		},
		{
			Name:       "otlphttp",
			Protocol:   "TCP",
			Port:       4318,
			TargetPort: intstr.FromInt(4318),
		},
		{
			Name: "metrics",
			Port: 8888,
		},
	}

	svc.Spec.Selector = commonLabels
	svc.Spec.ClusterIP = "None"
}
