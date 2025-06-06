package generated

import "github.com/bertinatto/testgrid/internal"

// This file is generated by go generate. DO NOT EDIT.

var Variants = map[string]internal.Variant{

	"periodic-ci-openshift-cluster-control-plane-machine-set-operator-release-4.20-periodics-e2e-aws": {
		Name:                "aws-control-plane-machine-set-operator",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-cluster-control-plane-machine-set-operator-release-4.20-periodics-e2e-azure": {
		Name:                "azure-control-plane-machine-set-operator",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-cluster-control-plane-machine-set-operator-release-4.20-periodics-e2e-gcp": {
		Name:                "gcp-control-plane-machine-set-operator",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-eng-ocp-qe-perfscale-ci-main-aws-4.20-nightly-x86-payload-control-plane-6nodes": {
		Name:                "qe-perfscale",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-hypershift-release-4.20-periodics-e2e-aks": {
		Name:                "hypershift-aks",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-hypershift-release-4.20-periodics-e2e-aws-ovn-conformance": {
		Name:                "aggregated-hypershift-ovn-conformance-4.20",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-hypershift-release-4.20-periodics-e2e-azure-kubevirt-ovn": {
		Name:                "hypershift-kubevirt-ovn-conformance-4.20",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-kni-eco-ci-cd-main-nightly-4.20-telcov10n-metal-single-node-spoke": {
		Name:                "telcov10n-metal-single-node-spoke",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-microshift-release-4.20-periodics-e2e-aws-ovn-ocp-conformance": {
		Name:                "microshift-ovn-conformance-parallel",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-microshift-release-4.20-periodics-e2e-aws-ovn-ocp-conformance-serial": {
		Name:                "microshift-ovn-conformance-serial",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-osde2e-main-nightly-4.20-osd-aws": {
		Name:                "openshift-dedicated-aws",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-osde2e-main-nightly-4.20-osd-gcp": {
		Name:                "openshift-dedicated-gcp-osde2e",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-osde2e-main-nightly-4.20-rosa-classic-sts": {
		Name:                "rosa-classic-sts-osde2e",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-aws-ovn": {
		Name:                "aws-ovn",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-aws-ovn-techpreview": {
		Name:                "aws-ovn-techpreview",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-aws-ovn-techpreview-serial": {
		Name:                "aws-ovn-techpreview-serial",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-aws-ovn-upgrade-out-of-change": {
		Name:                "aws-ovn-upgrade-out-of-change",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  true,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-aws-upgrade-ovn-single-node": {
		Name:                "aggregated-aws-ovn-single-node-upgrade-4.20-micro",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  true,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-azure-ovn": {
		Name:                "azure-ovn",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-azure-ovn-serial": {
		Name:                "azure-ovn-serial",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-azure-ovn-techpreview": {
		Name:                "azure-ovn-techpreview",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-azure-ovn-techpreview-serial": {
		Name:                "azure-ovn-techpreview-serial",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-azure-ovn-upgrade": {
		Name:                "aggregated-azure-ovn-upgrade-4.20-micro",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  true,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-azure-ovn-upgrade-out-of-change": {
		Name:                "azure-ovn-upgrade-out-of-change",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  true,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-gcp-ovn": {
		Name:                "gcp-ovn",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-gcp-ovn-techpreview": {
		Name:                "gcp-ovn-techpreview",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-gcp-ovn-techpreview-serial": {
		Name:                "gcp-ovn-techpreview-serial",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-gcp-ovn-upgrade": {
		Name:                "gcp-ovn-upgrade-micro",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  true,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-e2e-vsphere-ovn-upgrade": {
		Name:                "vsphere-ovn-upgrade-micro",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  true,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-upgrade-from-stable-4.19-e2e-aws-ovn-upgrade": {
		Name:                "upgrade-minor-ovn",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: true,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-upgrade-from-stable-4.19-e2e-azure-ovn-upgrade": {
		Name:                "aggregated-azure-ovn-upgrade-4.20-minor",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: true,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-upgrade-from-stable-4.19-e2e-gcp-ovn-rt-upgrade": {
		Name:                "aggregated-gcp-ovn-rt-upgrade-4.20-minor",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: true,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-upgrade-from-stable-4.19-e2e-gcp-ovn-upgrade": {
		Name:                "gcp-ovn-upgrade-4.20-minor",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: true,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-ci-4.20-upgrade-from-stable-4.19-e2e-vsphere-ovn-upgrade": {
		Name:                "vsphere-ovn-upgrade-minor",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: true,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-cnv-nightly-4.20-deploy-azure-kubevirt-ovn": {
		Name:                "cnv",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-console-aws": {
		Name:                "aws-console",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-agent-compact-fips": {
		Name:                "agent-compact-ipv4",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-agent-ha-dualstack-conformance": {
		Name:                "agent-ha-dualstack",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-agent-single-node-ipv6-conformance": {
		Name:                "agent-single-node-ipv6",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-aws-csi": {
		Name:                "aws-csi",
		Parallel:            true,
		CSI:                 true,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-aws-driver-toolkit": {
		Name:                "driver-toolkit",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-aws-ovn-cgroupsv2": {
		Name:                "aws-ovn-cgroupsv2",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-aws-ovn-fips": {
		Name:                "aws-ovn-fips",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-aws-ovn-proxy": {
		Name:                "ovn-proxy",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-aws-ovn-serial": {
		Name:                "aws-ovn-serial",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-aws-ovn-single-node": {
		Name:                "aws-ovn-single-node",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-aws-ovn-single-node-csi": {
		Name:                "aws-ovn-single-node-csi",
		Parallel:            true,
		CSI:                 true,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-aws-ovn-single-node-serial": {
		Name:                "aws-ovn-single-node-serial",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-aws-ovn-single-node-techpreview": {
		Name:                "aws-ovn-single-node-techpreview",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-aws-ovn-single-node-techpreview-serial": {
		Name:                "aws-ovn-single-node-techpreview-serial",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-aws-ovn-upgrade-fips": {
		Name:                "aggregated-aws-ovn-upgrade-4.20-micro-fips",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  true,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-aws-ovn-upi": {
		Name:                "aws-upi",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-azure-csi": {
		Name:                "azure-csi",
		Parallel:            true,
		CSI:                 true,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-gcp-ovn-csi": {
		Name:                "gcp-ovn-csi",
		Parallel:            true,
		CSI:                 true,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-gcp-ovn-rt": {
		Name:                "gcp-ovn-rt",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-gcp-ovn-serial": {
		Name:                "gcp-ovn-serial",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-metal-ipi-ovn-bm": {
		Name:                "metal-ipi-ovn-bm",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-metal-ipi-ovn-bm-upgrade": {
		Name:                "metal-ipi-ovn-bm-upgrade",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  true,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-metal-ipi-ovn-dualstack": {
		Name:                "metal-ipi-ovn-dualstack",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-metal-ipi-ovn-dualstack-techpreview": {
		Name:                "metal-ipi-ovn-dualstack-techpreview",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-metal-ipi-ovn-ipv6": {
		Name:                "metal-ipi-ovn-ipv6",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-metal-ipi-ovn-ipv6-techpreview": {
		Name:                "metal-ipi-ovn-ipv6-techpreview",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-metal-ipi-ovn-serial-ipv4": {
		Name:                "metal-ipi-ovn-serial",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-metal-ipi-ovn-serial-virtualmedia": {
		Name:                "metal-ipi-ovn-serial-virtualmedia",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-metal-ipi-ovn-techpreview": {
		Name:                "metal-ipi-ovn-techpreview",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-metal-ipi-serial-ovn-dualstack": {
		Name:                "metal-ipi-serial-ovn-dualstack",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-metal-ipi-serial-ovn-ipv6": {
		Name:                "metal-ipi-serial-ipv6",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-metal-ipi-upgrade-ovn-ipv6": {
		Name:                "metal-ipi-upgrade-ovn-ipv6",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  true,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-metal-ovn-single-node-live-iso": {
		Name:                "ovn-single-node-live-iso",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-osd-ccs-gcp": {
		Name:                "openshift-dedicated-gcp-conformance",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-rosa-sts-hypershift-ovn": {
		Name:                "rosa-hypershift-sts-conformance",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-rosa-sts-ovn": {
		Name:                "rosa-classic-sts-conformance",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-telco5g": {
		Name:                "telco5g",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-vsphere-ovn": {
		Name:                "vsphere-ovn",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-vsphere-ovn-csi": {
		Name:                "vsphere-ovn-csi",
		Parallel:            true,
		CSI:                 true,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-vsphere-ovn-serial": {
		Name:                "vsphere-ovn-serial",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-vsphere-ovn-techpreview": {
		Name:                "vsphere-ovn-techpreview",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-vsphere-ovn-techpreview-serial": {
		Name:                "vsphere-ovn-techpreview-serial",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-vsphere-ovn-upi": {
		Name:                "vsphere-ovn-upi",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-vsphere-ovn-upi-serial": {
		Name:                "vsphere-ovn-upi-serial",
		Parallel:            false,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              true,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-e2e-vsphere-static-ovn": {
		Name:                "vsphere-static-ovn",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-fips-payload-scan": {
		Name:                "fips-scan",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-install-analysis-all": {
		Name:                "install-analysis-all",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-metal-ovn-single-node-recert-cluster-rename": {
		Name:                "metal-ovn-single-node-recert-cluster-rename",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-overall-analysis-all": {
		Name:                "overall-analysis-all",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-upgrade-analysis-all": {
		Name:                "upgrade-analysis-all",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: false,
		UpgradeFromCurrent:  true,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-upgrade-from-stable-4.19-e2e-aws-upgrade-ovn-single-node": {
		Name:                "aws-ovn-single-node-upgrade-minor",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: true,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-upgrade-from-stable-4.19-e2e-metal-ipi-ovn-upgrade": {
		Name:                "metal-ipi-ovn-upgrade-minor",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: true,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
	"periodic-ci-openshift-release-master-nightly-4.20-upgrade-from-stable-4.19-e2e-metal-ipi-upgrade-ovn-ipv6": {
		Name:                "metal-ipi-upgrade-ovn-ipv6-minor",
		Parallel:            true,
		CSI:                 false,
		UpgradeFromPrevious: true,
		UpgradeFromCurrent:  false,
		Serial:              false,
	},
}
