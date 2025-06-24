# Argus service

## API Endpoints

> GET https://argus.service.glaucus.net/orders
```json
[
    {
        "id": "aa79b94d-4753-48f9-8eef-0dc1314e49ff",
        "url": "https://nurio.me/cv.pdf",
        "checkType": "Status:200",
        "notify": {
            "phone": "+34666333555",
            "title": "NurioMe",
            "message": "Nurio CV is online!"
        },
        "period": "1h"
    }
]
```

> POST https://argus.service.glaucus.net/orders/add
```json
{
    "url": "https://nurio.me/cv.pdf",
    "checkType": "Status:200",
    "notify": {
        "phone": "+34666333555",
        "title": "NurioMe",
        "message": "Nurio CV is online!"
    },
    "period": "1h"
}
```
