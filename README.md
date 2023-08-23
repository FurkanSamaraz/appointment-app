# appointment-app

Bu proje, müşteri randevularını yönetmek için bir uygulama geliştirmek amacıyla oluşturulmuştur. Uygulama, müşteri bilgilerini, randevu tarihlerini ve verilen hizmetleri içeren bir veritabanını yönetir.

## Özellikler

- Müşteri yönetimi: Müşterilerin adı, soyadı, e-posta ve telefon numarası bilgileri kaydedilebilir.
- Randevu yönetimi: Müşterilere ait randevu tarihleri ve amaçları kaydedilir.
- Hizmet yönetimi: Sunulan hizmetlerin adı, açıklaması ve fiyatı kaydedilir.
- Müşteri hizmet geçmişi: Müşterilere verilen hizmetler ve randevu tarihleri kaydedilir.

## Kullanım

1. Uygulamayı çalıştırmadan önce ".env" dosyasını doldurun. Bu dosya, PostgreSQL veritabanı ve uygulama ayarlarını içerir.
2. Veritabanı tablolarını oluşturmak için "configs/postgres/migrations/create_table_up.sql" dosyasını kullanın veya otomatik migrasyonu kullanın (opsiyonel).
3. Uygulamayı çalıştırın: `go run main.go`.

## Bağımlılıklar

- Golang (Go) 1.15 veya üstü
- PostgreSQL veritabanı
- minikube (k8s)
- docker

## Kurulum

1. Bu depoyu klonlayın: `git clone <repo-url>`
2. Proje klasörüne gidin: `cd appointment-app`
3. Gerekli bağımlılıkları yükleyin: `go mod tidy`
