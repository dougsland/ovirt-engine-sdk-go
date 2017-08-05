package ovirtsdk4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXMLClusterReadOne(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<cluster href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310" id="00000002-0002-0002-0002-000000000310">
    <actions>
        <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/resetemulatedmachine" rel="resetemulatedmachine"/>
    </actions>
    <name>Default</name>
    <description>The default server cluster</description>
    <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/networks" rel="networks"/>
    <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/permissions" rel="permissions"/>
    <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/glustervolumes" rel="glustervolumes"/>
    <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/glusterhooks" rel="glusterhooks"/>
    <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/affinitygroups" rel="affinitygroups"/>
    <link href="/ovirt-engine/api/clusters/00000002-0002-0002-0002-000000000310/cpuprofiles" rel="cpuprofiles"/>
    <ballooning_enabled>false</ballooning_enabled>
    <cpu>
        <architecture>x86_64</architecture>
        <type>Intel SandyBridge Family</type>
    </cpu>
    <error_handling>
        <on_error>migrate</on_error>
    </error_handling>
    <fencing_policy>
        <enabled>true</enabled>
        <skip_if_connectivity_broken>
            <enabled>false</enabled>
            <threshold>50</threshold>
        </skip_if_connectivity_broken>
        <skip_if_sd_active>
            <enabled>false</enabled>
        </skip_if_sd_active>
    </fencing_policy>
    <gluster_service>false</gluster_service>
    <ha_reservation>false</ha_reservation>
    <ksm>
        <enabled>true</enabled>
        <merge_across_nodes>true</merge_across_nodes>
    </ksm>
    <maintenance_reason_required>false</maintenance_reason_required>
    <memory_policy>
        <over_commit>
            <percent>100</percent>
        </over_commit>
        <transparent_hugepages>
            <enabled>true</enabled>
        </transparent_hugepages>
    </memory_policy>
    <migration>
        <auto_converge>inherit</auto_converge>
        <bandwidth>
            <assignment_method>auto</assignment_method>
        </bandwidth>
        <compressed>inherit</compressed>
        <policy id="80554327-0569-496b-bdeb-fcbbf52b827b"/>
    </migration>
    <optional_reason>false</optional_reason>
    <required_rng_sources>
		<required_rng_source>random</required_rng_source>
		<required_rng_source>get</required_rng_source>
		<required_rng_source>post</required_rng_source>
    </required_rng_sources>
    <switch_type>legacy</switch_type>
    <threads_as_cores>false</threads_as_cores>
    <trusted_service>false</trusted_service>
    <tunnel_migration>false</tunnel_migration>
    <version>
        <major>4</major>
        <minor>0</minor>
    </version>
    <virt_service>true</virt_service>
    <data_center href="/ovirt-engine/api/datacenters/00000001-0001-0001-0001-0000000002ed" id="00000001-0001-0001-0001-0000000002ed"/>
    <scheduling_policy href="/ovirt-engine/api/schedulingpolicies/b4ed2332-a7ac-4d5f-9596-99a439cb2812" id="b4ed2332-a7ac-4d5f-9596-99a439cb2812"/>
</cluster>
`

	reader := NewXMLReader([]byte(xmlstring))

	cluster, err := XMLClusterReadOne(reader, nil)
	assert.Nil(err)
	// Cluster>Id
	// assert.Equal("00000002-0002-0002-0002-000000000310", *cluster.Id)
	// Cluster>Name
	assert.Equal("Default", *cluster.Name, "Name should be `Default`")
	// Cluster>CPU>Architecture
	assert.NotNil(cluster.Cpu.Architecture)
	assert.Equal(Architecture("x86_64"), *cluster.Cpu.Architecture, "CPU Arch should be `x86_64`")
	// Cluster>BallooningEnabled
	assert.False(*cluster.BallooningEnabled, "Cluster>BallooningEnabled should be false")
	// Cluster>Description
	assert.Equal("The default server cluster", *cluster.Description)
	// Cluster>ErrorHandling>OnError>MigrateOnError
	assert.NotNil(cluster.ErrorHandling, "Cluster>ErrorHandling should not be nil")
	assert.NotNil(cluster.ErrorHandling.OnError)
	assert.Equal(MigrateOnError("migrate"), *cluster.ErrorHandling.OnError)
	// Cluster>FencingPolicy
	assert.NotNil(*cluster.FencingPolicy)
	// 		>Enable
	assert.True(*cluster.FencingPolicy.Enabled, "Cluster>FencingPolicy>Enable should be true")
	// 		>SkipIfConnectivityBroken>Threshold
	assert.Equal(int64(50), *cluster.FencingPolicy.SkipIfConnectivityBroken.Threshold,
		"Cluster>FencingPolicy>SkipIfConnectivityBroken>Threshold should be 50")
	assert.False(*cluster.FencingPolicy.SkipIfSdActive.Enabled, "Cluster>FencingPolicy>SkipIfSdActive>Enabled should be false")
	// Cluster>SchedulingPolicy
	assert.Equal("b4ed2332-a7ac-4d5f-9596-99a439cb2812", *cluster.SchedulingPolicy.Id)
	assert.Equal("/ovirt-engine/api/schedulingpolicies/b4ed2332-a7ac-4d5f-9596-99a439cb2812", *cluster.SchedulingPolicy.Href)
}

func TestXMLErrorHandlingReadOne(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `
    <error_handling>
        <on_error>migrate</on_error>
	</error_handling>`
	reader := NewXMLReader([]byte(xmlstring))
	eh, err := XMLErrorHandlingReadOne(reader, nil)
	assert.Nil(err)
	assert.NotNil(eh)
	assert.NotNil(eh.OnError)
	assert.Equal(MigrateOnError("migrate"), *eh.OnError)
}

func TestXMLCPUReadOne(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<cpu>
	<architecture>x86_64</architecture>
	<type>Intel SandyBridge Family</type>
</cpu>
`

	reader := NewXMLReader([]byte(xmlstring))

	cpu, err := XMLCpuReadOne(reader, nil)
	assert.Nil(err)
	assert.NotNil(cpu.Architecture)
	assert.Equal(Architecture("x86_64"), *cpu.Architecture)
	assert.Equal("Intel SandyBridge Family", *cpu.Type)
}

func TestArchitectureReadMany(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<architectures>
	<architecture>x86_64</architecture>
	<architecture>x86</architecture>
	<architecture>mips</architecture>
	<architecture>powerpc</architecture>
</architectures>
`

	reader := NewXMLReader([]byte(xmlstring))

	archs, err := XMLArchitectureReadMany(reader, nil)
	assert.Nil(err)
	assert.Equal([]Architecture{
		Architecture("x86_64"), Architecture("x86"),
		Architecture("mips"), Architecture("powerpc")}, archs)
}

func TestFaultReadOne(t *testing.T) {
	assert := assert.New(t)
	xmlstring := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<architectures>
	<architecture>x86_64</architecture>
	<architecture>x86</architecture>
	<architecture>mips</architecture>
	<architecture>powerpc</architecture>
</architectures>`
	reader := NewXMLReader([]byte(xmlstring))
	fault, err := XMLFaultReadOne(reader, nil)
	assert.Nil(err)
	assert.Nil(fault)
}
