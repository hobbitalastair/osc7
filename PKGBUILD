# Maintainer: ash
pkgname=osc7
pkgver=1
pkgrel=1
pkgdesc='Print an OSC 7 escape sequence for the current working directory'
arch=('x86_64' 'aarch64')
url=''
license=('MIT')
depends=('glibc')
makedepends=('git' 'go')
source=("$pkgname-src::git+file://$PWD#branch=main")
sha256sums=('SKIP')

build() {
	cd "$srcdir/$pkgname-src"
	make
}

check() {
	cd "$srcdir/$pkgname-src"
	make test
}

package() {
	cd "$srcdir/$pkgname-src"
	make DESTDIR="$pkgdir" PREFIX=/usr install
	install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
}
