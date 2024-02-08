# fergo - šatrovački Gopher server 

*fergo* je Gopher server implementiran u programskom jeziku Go i moj maturski rad u Drugoj gimnaziji. *fergo* može klijentu ispisati sadržaj odabrane datoteke ili direktorija (u obliku menija sa linkovima na druge direktorije i datoteke), a administratoru omogućuje da pravi vlastite menije prateći *user-friendly* sintaksu (kao [geomyidae](http://r-36.net/scm/geomyidae/file/README.html)). Korisnik također može odrediti i *hostname* prikazan u menijima, mrežni interfejs i TCP port koje "sluša" server, direktorij iz kojeg server poslužuje datoteke, datoteku u koji se bilježe *log*-ovi te verziju internet protokola (IPv4 ili IPv6).

## Instalacija

`go install github.com/amuradbegovic/fergo@latest`  