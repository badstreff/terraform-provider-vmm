[![CodeFactor](https://www.codefactor.io/repository/github/badstreff/terraform-provider-vmm/badge)](https://www.codefactor.io/repository/github/badstreff/terraform-provider-vmm)
[![codecov](https://codecov.io/gh/badstreff/terraform-provider-vmm/branch/master/graph/badge.svg)](https://codecov.io/gh/badstreff/terraform-provider-vmm)
[![Build Status](https://travis-ci.org/badstreff/terraform-provider-vmm.svg?branch=master)](https://travis-ci.org/badstreff/terraform-provider-vmm)

Terraform Provider VMM
---

Terraform provider for Microsoft's System Center Virtual Machine Manager product that leverages the [Service Provider Foundation SDK](https://docs.microsoft.com/en-us/previous-versions/system-center/developer/jj643273(v%3dmsdn.10)) for talking to VMM. This project differs from the SCVMM provider as it does not leverage any powershell scripts to copy to the hosts for execution, instead it relies on the SPF SDK for doing the heavy lifting on the backend.

Usage
---

Refer to the examples directory for the resources available from this provider.


Contributing
---

PRs Welcome! Open an issue first if you have concerns.

