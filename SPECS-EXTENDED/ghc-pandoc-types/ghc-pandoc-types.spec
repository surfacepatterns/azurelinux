Vendor:         Microsoft Corporation
Distribution:   Mariner
# generated by cabal-rpm-2.0.9
# https://docs.fedoraproject.org/en-US/packaging-guidelines/Haskell/

%global pkg_name pandoc-types
%global pkgver %{pkg_name}-%{version}

# testsuite missing deps: test-framework test-framework-hunit test-framework-quickcheck2

Name:           ghc-%{pkg_name}
Version:        1.22.1
Release:        2%{?dist}
Summary:        Types for representing a structured document

License:        BSD
Url:            https://hackage.haskell.org/package/%{pkg_name}
# Begin cabal-rpm sources:
Source0:        https://hackage.haskell.org/package/%{pkgver}/%{pkgver}.tar.gz
# End cabal-rpm sources

# Begin cabal-rpm deps:
BuildRequires:  ghc-Cabal-devel
BuildRequires:  ghc-rpm-macros
BuildRequires:  ghc-QuickCheck-prof
BuildRequires:  ghc-aeson-prof
BuildRequires:  ghc-base-prof
BuildRequires:  ghc-bytestring-prof
BuildRequires:  ghc-containers-prof
BuildRequires:  ghc-deepseq-prof
BuildRequires:  ghc-syb-prof
BuildRequires:  ghc-text-prof
BuildRequires:  ghc-transformers-prof
# End cabal-rpm deps

%description
This package contains definitions for the Pandoc data structure, which
is used by pandoc to represent structured documents. These definitions
used to live in the pandoc package, but they have been split off, so
that other packages can use them without drawing in all of pandoc's
dependencies, and pandoc itself can depend on packages (like
citeproc-hs) that use them.


%package devel
Summary:        Haskell %{pkg_name} library development files
Provides:       %{name}-static = %{version}-%{release}
Provides:       %{name}-static%{?_isa} = %{version}-%{release}
%if %{defined ghc_version}
Requires:       ghc-compiler = %{ghc_version}
%endif
Requires:       %{name}%{?_isa} = %{version}-%{release}

%description devel
This package provides the Haskell %{pkg_name} library development files.


%if %{with haddock}
%package doc
Summary:        Haskell %{pkg_name} library documentation
BuildArch:      noarch
Requires:       ghc-filesystem

%description doc
This package provides the Haskell %{pkg_name} library documentation.
%endif


%if %{with ghc_prof}
%package prof
Summary:        Haskell %{pkg_name} profiling library
Requires:       %{name}-devel%{?_isa} = %{version}-%{release}
Supplements:    (%{name}-devel and ghc-prof)

%description prof
This package provides the Haskell %{pkg_name} profiling library.
%endif


%prep
# Begin cabal-rpm setup:
%setup -q -n %{pkgver}
# End cabal-rpm setup


%build
# Begin cabal-rpm build:
%ghc_lib_build
# End cabal-rpm build


%install
# Begin cabal-rpm install
%ghc_lib_install
# End cabal-rpm install


%files -f %{name}.files
# Begin cabal-rpm files:
%license LICENSE
# End cabal-rpm files


%files devel -f %{name}-devel.files
%doc changelog


%if %{with haddock}
%files doc -f %{name}-doc.files
%license LICENSE
%endif


%if %{with ghc_prof}
%files prof -f %{name}-prof.files
%endif


%changelog
* Thu Jul 21 2022 Fedora Release Engineering <releng@fedoraproject.org> - 1.22.1-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_37_Mass_Rebuild

* Tue Jun 07 2022 Jens Petersen <petersen@redhat.com> - 1.22.1-1
- https://hackage.haskell.org/package/pandoc-types-1.22.1/changelog

* Thu Jan 20 2022 Fedora Release Engineering <releng@fedoraproject.org> - 1.22-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_36_Mass_Rebuild

* Thu Aug  5 2021 Jens Petersen <petersen@redhat.com> - 1.22-1
- update to 1.22

* Thu Jul 22 2021 Fedora Release Engineering <releng@fedoraproject.org> - 1.20-5
- Rebuilt for https://fedoraproject.org/wiki/Fedora_35_Mass_Rebuild

* Tue Jan 26 2021 Fedora Release Engineering <releng@fedoraproject.org> - 1.20-4
- Rebuilt for https://fedoraproject.org/wiki/Fedora_34_Mass_Rebuild

* Sat Aug 01 2020 Fedora Release Engineering <releng@fedoraproject.org> - 1.20-3
- Second attempt - Rebuilt for
  https://fedoraproject.org/wiki/Fedora_33_Mass_Rebuild

* Mon Jul 27 2020 Fedora Release Engineering <releng@fedoraproject.org> - 1.20-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_33_Mass_Rebuild

* Wed Jun 10 2020 Jens Petersen <petersen@redhat.com> - 1.20-1
- update to 1.20

* Fri Feb 14 2020 Jens Petersen <petersen@redhat.com> - 1.17.6.1-1
- update to 1.17.6.1

* Tue Jan 28 2020 Fedora Release Engineering <releng@fedoraproject.org> - 1.17.5.4-4
- Rebuilt for https://fedoraproject.org/wiki/Fedora_32_Mass_Rebuild

* Fri Aug 02 2019 Jens Petersen <petersen@redhat.com> - 1.17.5.4-3
- add doc and prof subpackages (cabal-rpm-1.0.0)

* Thu Jul 25 2019 Fedora Release Engineering <releng@fedoraproject.org> - 1.17.5.4-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_31_Mass_Rebuild

* Thu Feb 21 2019 Jens Petersen <petersen@redhat.com> - 1.17.5.4-1
- update to 1.17.5.4

* Sun Feb 17 2019 Jens Petersen <petersen@redhat.com> - 1.17.3.1-3
- refresh to cabal-rpm-0.13

* Thu Jan 31 2019 Fedora Release Engineering <releng@fedoraproject.org> - 1.17.3.1-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_30_Mass_Rebuild

* Wed Aug  8 2018 Jens Petersen <petersen@redhat.com>
- update License tag to BSD

* Sat Jul 28 2018 Jens Petersen <petersen@redhat.com> - 1.17.3.1-1
- update to 1.17.3.1

* Mon Jul 23 2018 Miro Hrončok <mhroncok@redhat.com> - 1.17.3-4
- Rebuilt for #1607054

* Fri Jul 13 2018 Fedora Release Engineering <releng@fedoraproject.org> - 1.17.3-3
- Rebuilt for https://fedoraproject.org/wiki/Fedora_29_Mass_Rebuild

* Wed Feb 07 2018 Fedora Release Engineering <releng@fedoraproject.org> - 1.17.3-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_28_Mass_Rebuild

* Wed Jan 24 2018 Jens Petersen <petersen@redhat.com> - 1.17.3-1
- update to 1.17.3

* Wed Aug 02 2017 Fedora Release Engineering <releng@fedoraproject.org> - 1.17.0.5-3
- Rebuilt for https://fedoraproject.org/wiki/Fedora_27_Binutils_Mass_Rebuild

* Wed Jul 26 2017 Fedora Release Engineering <releng@fedoraproject.org> - 1.17.0.5-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_27_Mass_Rebuild

* Wed Feb 22 2017 Jens Petersen <petersen@redhat.com> - 1.17.0.5-1
- update to 1.17.0.5

* Fri Feb 10 2017 Fedora Release Engineering <releng@fedoraproject.org> - 1.16.1-3
- Rebuilt for https://fedoraproject.org/wiki/Fedora_26_Mass_Rebuild

* Fri Jul  1 2016 Jens Petersen <petersen@redhat.com> - 1.16.1-2
- add changelog

* Sat Mar 05 2016 Jens Petersen <petersen@redhat.com> - 1.16.1-1
- update to 1.16.1

* Wed Feb 03 2016 Fedora Release Engineering <releng@fedoraproject.org> - 1.12.4.1-4
- Rebuilt for https://fedoraproject.org/wiki/Fedora_24_Mass_Rebuild

* Mon Aug 31 2015 Peter Robinson <pbrobinson@fedoraproject.org> 1.12.4.1-3
- Rebuild (aarch64 vector hashes)

* Wed Jun 17 2015 Fedora Release Engineering <rel-eng@lists.fedoraproject.org> - 1.12.4.1-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_23_Mass_Rebuild

* Mon Jan 26 2015 Jens Petersen <petersen@redhat.com> - 1.12.4.1-1
- update to 1.12.4.1

* Sat Aug 16 2014 Fedora Release Engineering <rel-eng@lists.fedoraproject.org> - 1.12.3.3-3
- Rebuilt for https://fedoraproject.org/wiki/Fedora_21_22_Mass_Rebuild

* Sat Jun 07 2014 Fedora Release Engineering <rel-eng@lists.fedoraproject.org> - 1.12.3.3-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_21_Mass_Rebuild

* Thu May 08 2014 Jens Petersen <petersen@redhat.com> - 1.12.3.3-1
- update to 1.12.3.3

* Sun Apr 20 2014 Jens Petersen <petersen@redhat.com> - 1.12.3.1-3
- revert requiring ghci: aeson now patched to build when no ghci

* Sun Apr 20 2014 Jens Petersen <petersen@redhat.com> - 1.12.3.1-2
- aeson needs TemplateHaskell
- update packaging to latest cblrpm

* Wed Jan 22 2014 Jens Petersen <petersen@redhat.com> - 1.12.3.1-1
- update to 1.12.3.1

* Sat Aug 03 2013 Fedora Release Engineering <rel-eng@lists.fedoraproject.org> - 1.10-3
- Rebuilt for https://fedoraproject.org/wiki/Fedora_20_Mass_Rebuild

* Wed Jun 05 2013 Jens Petersen <petersen@redhat.com> - 1.10-2
- update to new simplified Haskell Packaging Guidelines

* Sun Mar 10 2013 Jens Petersen <petersen@redhat.com> - 1.10-1
- update to 1.10

* Wed Feb 13 2013 Fedora Release Engineering <rel-eng@lists.fedoraproject.org> - 1.9.1-7
- Rebuilt for https://fedoraproject.org/wiki/Fedora_19_Mass_Rebuild

* Sat Nov 17 2012 Jens Petersen <petersen@redhat.com> - 1.9.1-6
- update with cabal-rpm

* Thu Jul 19 2012 Fedora Release Engineering <rel-eng@lists.fedoraproject.org> - 1.9.1-5
- Rebuilt for https://fedoraproject.org/wiki/Fedora_18_Mass_Rebuild

* Mon Jul 16 2012 Jens Petersen <petersen@redhat.com> - 1.9.1-4
- change prof BRs to devel

* Fri Jun 15 2012 Jens Petersen <petersen@redhat.com> - 1.9.1-3
- rebuild

* Thu Mar 22 2012 Jens Petersen <petersen@redhat.com> - 1.9.1-2
- add license to ghc_files

* Wed Mar  7 2012 Jens Petersen <petersen@redhat.com> - 1.9.1-1
- update to 1.9.1

* Sun Feb 12 2012 Jens Petersen <petersen@redhat.com> - 1.9.0.2-1
- update to 1.9.0.2

* Thu Jan  5 2012 Jens Petersen <petersen@redhat.com> - 1.8.2-2
- update to cabal2spec-0.25.2

* Mon Oct 24 2011 Marcela Mašláňová <mmaslano@redhat.com> - 1.8.2-1.3
- rebuild with new gmp without compat lib

* Fri Oct 21 2011 Marcela Mašláňová <mmaslano@redhat.com> - 1.8.2-1.2
- rebuild with new gmp without compat lib

* Tue Oct 11 2011 Peter Schiffer <pschiffe@redhat.com> - 1.8.2-1.1
- rebuild with new gmp

* Thu Aug  4 2011 Jens Petersen <petersen@redhat.com> - 1.8.2-1
- update to 1.8.2

* Wed Jun 22 2011 Jens Petersen <petersen@redhat.com> - 1.8.0.2-2
- use ghc_arches (cabal2spec-0.23.2)

* Fri May 13 2011 Jens Petersen <petersen@redhat.com> - 1.8.0.2-1
- update to 1.8.0.2

* Wed May 11 2011 Jens Petersen <petersen@redhat.com> - 1.8-2
- drop ghc_package_deps
- add BR on ghc-Cabal-devel and ghc-containers-prof

* Tue May  3 2011 Jens Petersen <petersen@redhat.com> - 1.8-1
- GPLv2+ and depends on syb

* Tue May  3 2011 Fedora Haskell SIG <haskell-devel@lists.fedoraproject.org> - 1.8-0
- initial packaging for Fedora automatically generated by cabal2spec-0.22.6
