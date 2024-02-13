# Enable PPIx::Regexp optional feature
%bcond_without perl_PPIx_QuoteLike_enables_PPIx_Regexp

Name:           perl-PPIx-QuoteLike
Version:        0.008
Release:        3%{?dist}
Summary:        Parse Perl string literals and string-literal-like things
License:        GPL+ or Artistic
Vendor:         Microsoft Corporation
Distribution:   Azure Linux
URL:            https://metacpan.org/release/PPIx-QuoteLike
Source0:        https://cpan.metacpan.org/authors/id/W/WY/WYANT/PPIx-QuoteLike-%{version}.tar.gz#/perl-PPIx-QuoteLike-%{version}.tar.gz
BuildArch:      noarch
BuildRequires:  make
BuildRequires:  perl-generators
BuildRequires:  perl-interpreter
# Build.PL and inc/My/Module/Build.pm not used
BuildRequires:  perl(:VERSION) >= 5.6
BuildRequires:  perl(Carp)
BuildRequires:  perl(Config)
BuildRequires:  perl(constant)
BuildRequires:  perl(Exporter)
BuildRequires:  perl(ExtUtils::MakeMaker) >= 6.76
BuildRequires:  perl(lib)
BuildRequires:  perl(strict)
# Test::Without::Module not helpful
BuildRequires:  perl(warnings)
# Run-time:
BuildRequires:  perl(base)
BuildRequires:  perl(Encode)
BuildRequires:  perl(List::Util)
BuildRequires:  perl(PPI::Document) >= 1.117
BuildRequires:  perl(PPI::Dumper) >= 1.117
BuildRequires:  perl(re)
BuildRequires:  perl(Scalar::Util)
%if %{with perl_PPIx_QuoteLike_enables_PPIx_Regexp}
# Optional run-time:
# Author states there is a build-cycle with PPIx::Regexp, but I cannot see
# any.
BuildRequires:  perl(PPIx::Regexp)
%endif
# Tests:
BuildRequires:  perl(charnames)
BuildRequires:  perl(open)
BuildRequires:  perl(Test::More) >= 0.88
Requires:       perl(:MODULE_COMPAT_%(eval "`perl -V:version`"; echo $version))
Requires:       perl(PPI::Document) >= 1.117
Requires:       perl(PPI::Dumper) >= 1.117
%if %{with perl_PPIx_QuoteLike_enables_PPIx_Regexp}
Recommends:     perl(PPIx::Regexp)
%endif

# Remove under-specified dependencies
%global __requires_exclude %{?__requires_exclude:%{__requires_exclude}|}^perl\\((PPI::Document|PPI::Dumper)\\)$

%description
This Perl class parses Perl string literals and things that are reasonably
like string literals. Its real reason for being is to find interpolated
variables for Perl::Critic policies and similar code.

%prep
%setup -q -n PPIx-QuoteLike-%{version}
# Fix shell bang and permissions
for F in eg/{pqldump,variables}; do
    perl -MConfig -p -i -e 's{\A#!/usr/bin/env perl\b}{$Config{startperl}}' \
        "$F"
    chmod -x "$F"
done

%build
perl Makefile.PL INSTALLDIRS=vendor NO_PACKLIST=1 NO_PERLLOCAL=1
%{make_build}

%install
%{make_install}
%{_fixperms} $RPM_BUILD_ROOT/*

%check
make test

%files
%license LICENSES/*
%doc Changes eg README
%{perl_vendorlib}/*
%{_mandir}/man3/*

%changelog
* Fri Oct 15 2021 Pawel Winogrodzki <pawelwi@microsoft.com> - 0.008-3
- Initial CBL-Mariner import from Fedora 32 (license: MIT).

* Thu Jan 30 2020 Fedora Release Engineering <releng@fedoraproject.org> - 0.008-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_32_Mass_Rebuild

* Mon Aug 19 2019 Petr Pisar <ppisar@redhat.com> - 0.008-1
- 0.008 bump

* Fri Jul 26 2019 Fedora Release Engineering <releng@fedoraproject.org> - 0.007-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_31_Mass_Rebuild

* Wed Jun 05 2019 Petr Pisar <ppisar@redhat.com> - 0.007-1
- 0.007 bump

* Fri May 31 2019 Jitka Plesnikova <jplesnik@redhat.com> - 0.006-4
- Perl 5.30 rebuild

* Fri Feb 01 2019 Fedora Release Engineering <releng@fedoraproject.org> - 0.006-3
- Rebuilt for https://fedoraproject.org/wiki/Fedora_30_Mass_Rebuild

* Fri Jul 13 2018 Fedora Release Engineering <releng@fedoraproject.org> - 0.006-2
- Rebuilt for https://fedoraproject.org/wiki/Fedora_29_Mass_Rebuild

* Tue Jul 10 2018 Petr Pisar <ppisar@redhat.com> - 0.006-1
- 0.006 bump

* Fri Jun 29 2018 Jitka Plesnikova <jplesnik@redhat.com> - 0.005-2
- Perl 5.28 rebuild

* Mon Jun 04 2018 Petr Pisar <ppisar@redhat.com> 0.005-1
- Specfile autogenerated by cpanspec 1.78.
