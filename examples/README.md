Below are 2 examples for creating a drive and another example for creating a VM from a template.

```
provider "vmm" {
  base_url = "https://example.com:8090/SC2012R2/VMM/Microsoft.Management.Odata.svc/"  # SC2016 should work as well
  username = "My-SPF-VMM-User"
  password = "secret-password"
  domain   = "example.com"
  stamp_id = "My-Unique-Stamp-Id"
}

resource "vmm_virtual_disk_drive" "drive" {
  vm_name = "MyNewVM"
  bus = 0
  file_name = "test_disk"
  lun = 1
  scsi = true
  stamp_id = "My-Unique-Stamp-Id"
  hard_disk_name = "Blank Disk - Small.vhdx"
}

resource "vmm_virtual_machine" "vm" {
  vm_name = "MyNewVM"
  stamp_id = "My-Unique-Stamp-Id"
  cloud_name = "Production"
  template_name = "ha-rhel-7"

  cpu_count = 4  # we can add more cpu/memory depending on the template
  
  # we can also adjust the network adapters on the template
  network_adapters = [
    {
      network_name = "TestDev"
      mac_type = "Static"
    }
  ]
}
```
