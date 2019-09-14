# Maintainer: Brenden Davidson <davidson.brenden15@gmail.com>
# Contributor: Brenden Davidson <davidson.brenden15@gmail.com>

buildarch=''
gover=1.13
gochksum=''

case $CARCH in
    x86_64)
        buildarch='linux-amd64'
        gochksum='68a2297eb099d1a76097905a2ce334e3155004ec08cdea85f24527be3c48e856'
        ;;
    i686)
        buildarch='linux-386'
        gochksum='519b3e6ae4db011b93b60e6fabb055773ae6448355b6909a6befef87e02d98f5'
        ;;
    armv6h)
        buildarch='linux-armv6l'
        gochksum='931906d67cae1222f501e7be26e0ee73ba89420be0c4591925901cb9a4e156f0'
        ;;
esac

pkgbase=zipproject
pkgname=('zipproject')
pkgver=v0.1.alpha.1.d5f8828
pkgrel=1
pkgdesc='Small program for intelligently zipping programming projects'
arch=('x86_64' 'i686' 'armv6h')
license=('MIT')

makedepends=('tar')
source=(
    "go$gover.$buildarch.tar.gz::https://dl.google.com/go/go$gover.$buildarch.tar.gz"
    "zipproject::git+https://github.com/iAmSomeone2/zip_project.git"
    )
noextract=("go$gover.$buildarch.tar.gz")
sha256sums=(
    $gochksum
    'SKIP'
    )

pkgver() {
  cd zipproject

  git describe --tags | sed 's/-/./; s/-g/./; s/-/./'
}

prepare() {
    tar -xzf go$gover.$buildarch.tar.gz
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
