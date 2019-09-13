# Maintainer: Brenden Davidson <davidson.brenden15@gmail.com>
# Contributor: Brenden Davidson <davidson.brenden15@gmail.com>

pkgbase=zipproject
pkgname=('zipproject')
pkgver=v0.1_alpha
pkgrel=1
pkgdesc='Small program for intelligently zipping programming projects'
arch=('x86_64')
license=('MIT')

makedepends=('tar')
source=(
    "go1.13.linux-amd64.tar.gz::https://dl.google.com/go/go1.13.linux-amd64.tar.gz"
    "zipproject::git+https://github.com/iAmSomeone2/zip_project.git"
    )
noextract=("go1.13.linux-amd64.tar.gz")
sha256sums=(
    '68a2297eb099d1a76097905a2ce334e3155004ec08cdea85f24527be3c48e856'
    'SKIP'
    )

pkgver() {
  cd zipproject

  git describe --tags | sed 's/-/_/'
}

prepare() {
    tar -xzf go1.13.linux-amd64.tar.gz
}

build() {
    cd zipproject
    ../go/bin/go build
}

package_zipproject() {
    provides=('zipproject')

    mkdir -p $pkgdir/usr/bin
    install -m 755 zipproject/zipproject $pkgdir/usr/bin/zipproject
}