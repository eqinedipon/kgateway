package wellknown

import (
	apiextv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	gwv1 "sigs.k8s.io/gateway-api/apis/v1"
	gwv1a2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
	gwv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

const (
	// Group string for Gateway API resources
	GatewayGroup = gwv1.GroupName
	// XListenerSetGroup is the promoted ListenerSet API group.
	// TODO: Rename the exported XListenerSet* identifiers to ListenerSet* in a follow-up cleanup.
	// We intentionally keep the legacy symbol names in this bump PR even though ListenerSet
	// moved from apisx/v1alpha1 to gateway.networking.k8s.io/v1 in Gateway API v1.5.1.
	XListenerSetGroup = gwv1.GroupName

	// Kind strings
	ServiceKind          = "Service"
	ConfigMapKind        = "ConfigMap"
	SecretKind           = "Secret"
	HTTPRouteKind        = "HTTPRoute"
	TCPRouteKind         = "TCPRoute"
	TLSRouteKind         = "TLSRoute"
	GRPCRouteKind        = "GRPCRoute"
	GatewayKind          = "Gateway"
	GatewayClassKind     = "GatewayClass"
	ReferenceGrantKind   = "ReferenceGrant"
	BackendTLSPolicyKind = "BackendTLSPolicy"

	// XListenerSetKind is the promoted ListenerSet kind.
	// TODO: Rename the exported XListenerSet* identifiers to ListenerSet* in a follow-up cleanup.
	XListenerSetKind = "ListenerSet"

	// List Kind strings
	HTTPRouteListKind      = "HTTPRouteList"
	GatewayListKind        = "GatewayList"
	GatewayClassListKind   = "GatewayClassList"
	ReferenceGrantListKind = "ReferenceGrantList"

	// Gateway API CRD names
	TCPRouteCRDName = "tcproutes.gateway.networking.k8s.io"
)

var (
	GatewayGVK = schema.GroupVersionKind{
		Group:   GatewayGroup,
		Version: gwv1.GroupVersion.Version,
		Kind:    GatewayKind,
	}
	GatewayGVR = schema.GroupVersionResource{
		Group:    GatewayGroup,
		Version:  gwv1.GroupVersion.Version,
		Resource: "gateways",
	}
	GatewayClassGVK = schema.GroupVersionKind{
		Group:   GatewayGroup,
		Version: gwv1.GroupVersion.Version,
		Kind:    GatewayClassKind,
	}
	GatewayClassGVR = schema.GroupVersionResource{
		Group:    GatewayGroup,
		Version:  gwv1.GroupVersion.Version,
		Resource: "gatewayclasses",
	}
	HTTPRouteGVK = schema.GroupVersionKind{
		Group:   GatewayGroup,
		Version: gwv1.GroupVersion.Version,
		Kind:    HTTPRouteKind,
	}
	HTTPRouteGVR = schema.GroupVersionResource{
		Group:    GatewayGroup,
		Version:  gwv1.GroupVersion.Version,
		Resource: "httproutes",
	}
	TLSRouteGVK = schema.GroupVersionKind{
		Group:   GatewayGroup,
		Version: gwv1a2.GroupVersion.Version,
		Kind:    TLSRouteKind,
	}
	TLSRouteGVR = schema.GroupVersionResource{
		Group:    GatewayGroup,
		Version:  gwv1a2.GroupVersion.Version,
		Resource: "tlsroutes",
	}
	TCPRouteGVK = schema.GroupVersionKind{
		Group:   GatewayGroup,
		Version: gwv1a2.GroupVersion.Version,
		Kind:    TCPRouteKind,
	}
	TCPRouteGVR = schema.GroupVersionResource{
		Group:    GatewayGroup,
		Version:  gwv1a2.GroupVersion.Version,
		Resource: "tcproutes",
	}
	GRPCRouteGVK = schema.GroupVersionKind{
		Group:   GatewayGroup,
		Version: gwv1.GroupVersion.Version,
		Kind:    GRPCRouteKind,
	}
	GRPCRouteGVR = schema.GroupVersionResource{
		Group:    GatewayGroup,
		Version:  gwv1.GroupVersion.Version,
		Resource: "grpcroutes",
	}
	ReferenceGrantGVK = schema.GroupVersionKind{
		Group:   GatewayGroup,
		Version: gwv1b1.GroupVersion.Version,
		Kind:    ReferenceGrantKind,
	}
	ReferenceGrantGVR = schema.GroupVersionResource{
		Group:    GatewayGroup,
		Version:  gwv1b1.GroupVersion.Version,
		Resource: "referencegrants",
	}
	BackendTLSPolicyGVK = schema.GroupVersionKind{
		Group:   GatewayGroup,
		Version: gwv1.GroupVersion.Version,
		Kind:    BackendTLSPolicyKind,
	}
	BackendTLSPolicyGVR = schema.GroupVersionResource{
		Group:    GatewayGroup,
		Version:  gwv1.GroupVersion.Version,
		Resource: "backendtlspolicies",
	}

	TCPRouteCRD = apiextv1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: TCPRouteCRDName,
		},
	}

	// XListenerSetGVK is the promoted ListenerSet GVK.
	// TODO: Rename the exported XListenerSet* identifiers to ListenerSet* in a follow-up cleanup.
	XListenerSetGVK = schema.GroupVersionKind{
		Group:   XListenerSetGroup,
		Version: gwv1.GroupVersion.Version,
		Kind:    XListenerSetKind,
	}
	// XListenerSetGVR is the promoted ListenerSet GVR.
	// TODO: Rename the exported XListenerSet* identifiers to ListenerSet* in a follow-up cleanup.
	XListenerSetGVR = schema.GroupVersionResource{
		Group:    XListenerSetGroup,
		Version:  gwv1.GroupVersion.Version,
		Resource: "listenersets",
	}
)
