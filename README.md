TO-DO App REST API
Bu proje, kullanıcıların görevleri (TO-DO'lar) yönetmesini sağlayan bir REST API uygulamasıdır. API, görevlerin oluşturulması, güncellenmesi, silinmesi ve adımlarının takibini sağlar. Ayrıca kullanıcılar için JWT tabanlı kimlik doğrulama ve rol yönetimi içerir.
Özellikler
- Kullanıcı girişi (JWT tabanlı kimlik doğrulama)
- Görev oluşturma, güncelleme, silme (CRUD işlemleri)
- Görev adımlarını ekleyip düzenleme
- Soft delete (silinmiş görevler ve adımlar)
- Görev tamamlanma oranı hesaplama
- Yönetici ve kullanıcı rolleriyle erişim kontrolü
Kullanım
API Sonuçları
1. Kullanıcı Girişi (Login)
- **Endpoint:** `POST /login`
- **Body:**
{
  "username": "admin",
  "password": "admin123"
}
- **Başarı Durumu (200 OK):**
{
  "token": "your_jwt_token"
}
- **Hata Durumu (401 Unauthorized):**
{
  "error": "Invalid credentials"
}
2. Görevleri Listele (Get ToDos)
- **Endpoint:** `GET /api/todos`
- **Header:** `Authorization: Bearer <token>`
- **Başarı Durumu (200 OK):**
[
  {
    "id": "36557f4e-56ab-43df-bcec-1e84a98b5908",
    "name": "Projeyi tamamla",
    "created_at": "2025-05-04T18:39:12.7028268+03:00",
    "updated_at": "2025-05-04T18:42:16.2206453+03:00",
    "deleted_at": "2025-05-04T18:50:30.0102761+03:00",
    "username": "admin",
    "percent": 50,
    "steps": [ ... ]
  }
]
3. Görev Oluşturma (Create ToDo)
- **Endpoint:** `POST /api/todos/create`
- **Body:**
{
  "name": "Yeni Görev"
}
- **Başarı Durumu (201 Created):**
{
  "id": "new-id",
  "name": "Yeni Görev",
  "created_at": "2025-05-05T01:11:52.8056023+03:00",
  "updated_at": "2025-05-05T01:11:52.8056023+03:00",
  "username": "admin",
  "percent": 0
}
4. Görev Güncelleme (Update ToDo)
- **Endpoint:** `PUT /api/todos/update`
- **Body:**
{
  "id": "existing-id",
  "name": "Yeni Güncellenmiş Görev"
}
- **Başarı Durumu (200 OK):**
{
  "id": "existing-id",
  "name": "Yeni Güncellenmiş Görev",
  "created_at": "2025-05-04T18:39:12.7028268+03:00",
  "updated_at": "2025-05-05T01:11:52.8056023+03:00",
  "username": "admin",
  "percent": 0
}
5. Görev Silme (Delete ToDo)
- **Endpoint:** `DELETE /api/todos/delete/{id}`
- **Başarı Durumu (204 No Content):** Görev silinir.
6. Adım Oluşturma (Create Step)
- **Endpoint:** `POST /api/steps/create`
- **Body:**
{
  "todo_id": "existing-todo-id",
  "content": "Yeni Adım",
  "done": false
}
- **Başarı Durumu (201 Created):**
{
  "id": "step-id",
  "todo_id": "existing-todo-id",
  "content": "Yeni Adım",
  "done": false,
  "created_at": "2025-05-05T01:11:52.8056023+03:00",
  "updated_at": "2025-05-05T01:11:52.8056023+03:00"
}
7. Adım Güncelleme (Update Step)
- **Endpoint:** `PUT /api/steps/update`
- **Body:**
{
  "id": "existing-step-id",
  "content": "Güncellenmiş Adım",
  "done": true
}
- **Başarı Durumu (200 OK):**
{
  "id": "existing-step-id",
  "content": "Güncellenmiş Adım",
  "done": true,
  "updated_at": "2025-05-05T01:11:52.8056023+03:00"
}
8. Adım Silme (Delete Step)
- **Endpoint:** `DELETE /api/steps/delete/{id}`
- **Başarı Durumu (204 No Content):** Adım silinir.
              # Go sum dosyası
    
Geliştirme
Bağımlılıkları Yükleyin
Go mod kullanarak bağımlılıkları yüklemek için terminalde aşağıdaki komutu çalıştırın:
`go mod tidy`
Sunucuyu Başlatın
Uygulamayı başlatmak için:
`go run cmd/main.go`
Sunucu `http://localhost:8080`'da çalışacaktır.
.
