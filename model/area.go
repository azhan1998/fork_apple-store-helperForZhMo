package model

type Area struct {
    Title string
    Code string
    ProductsJson string
}

// Areas 地区，中国大陆: CN/zh_CN, 中国澳门: MO/zh_MO
var Areas = []Area {

    {
       Title: "中国澳门",
       Code: "MO/zh_MO",
       ProductsJson: `[{
            "partNumber": "MLT53ZA/A",
            "familyType": "iPhone 13 Pro 128GB 石墨色",
            "color": "石墨色",
            "capacity": "128GB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$8,799.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-graphite-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629910158000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-graphite-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629910158000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLHD3ZA/A",
            "familyType": "iPhone 13 Pro Max 512GB 石墨色",
            "color": "石墨色",
            "capacity": "512GB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$12,299.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-graphite-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629910158000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-graphite-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629910158000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLHG3ZA/A",
            "familyType": "iPhone 13 Pro Max 512GB 天峰藍色",
            "color": "天峰藍色",
            "capacity": "512GB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$12,299.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-blue-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841926000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-blue-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841926000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLTN3ZA/A",
            "familyType": "iPhone 13 Pro 1TB 天峰藍色",
            "color": "天峰藍色",
            "capacity": "1TB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$13,199.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-blue-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841925000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-blue-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841925000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLT83ZA/A",
            "familyType": "iPhone 13 Pro 128GB 天峰藍色",
            "color": "天峰藍色",
            "capacity": "128GB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$8,799.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-blue-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841925000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-blue-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841925000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLHC3ZA/A",
            "familyType": "iPhone 13 Pro Max 256GB 天峰藍色",
            "color": "天峰藍色",
            "capacity": "256GB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$10,499.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-blue-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841926000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-blue-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841926000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLTJ3ZA/A",
            "familyType": "iPhone 13 Pro 512GB 天峰藍色",
            "color": "天峰藍色",
            "capacity": "512GB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$11,399.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-blue-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841925000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-blue-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841925000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLT93ZA/A",
            "familyType": "iPhone 13 Pro 256GB 石墨色",
            "color": "石墨色",
            "capacity": "256GB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$9,699.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-graphite-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629910158000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-graphite-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629910158000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLTE3ZA/A",
            "familyType": "iPhone 13 Pro 256GB 天峰藍色",
            "color": "天峰藍色",
            "capacity": "256GB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$9,699.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-blue-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841925000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-blue-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841925000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLHF3ZA/A",
            "familyType": "iPhone 13 Pro Max 512GB 金色",
            "color": "gold",
            "capacity": "512GB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$12,299.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-gold-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841927000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-gold-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841927000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLTM3ZA/A",
            "familyType": "iPhone 13 Pro 1TB 金色",
            "color": "gold",
            "capacity": "1TB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$13,199.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-gold-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841924000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-gold-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841924000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLTK3ZA/A",
            "familyType": "iPhone 13 Pro 1TB 石墨色",
            "color": "石墨色",
            "capacity": "1TB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$13,199.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-graphite-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629910158000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-graphite-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629910158000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLH73ZA/A",
            "familyType": "iPhone 13 Pro Max 128GB 天峰藍色",
            "color": "天峰藍色",
            "capacity": "128GB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$9,699.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-blue-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841926000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-blue-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841926000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLH43ZA/A",
            "familyType": "iPhone 13 Pro Max 128GB 石墨色",
            "color": "石墨色",
            "capacity": "128GB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$9,699.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-graphite-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629910158000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-graphite-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629910158000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLTH3ZA/A",
            "familyType": "iPhone 13 Pro 512GB 金色",
            "color": "gold",
            "capacity": "512GB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$11,399.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-gold-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841924000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-gold-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841924000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLTC3ZA/A",
            "familyType": "iPhone 13 Pro 256GB 銀色",
            "color": "銀色",
            "capacity": "256GB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$9,699.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-silver-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841924000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-silver-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841924000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLT63ZA/A",
            "familyType": "iPhone 13 Pro 128GB 銀色",
            "color": "銀色",
            "capacity": "128GB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$8,799.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-silver-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841924000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-silver-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841924000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLHA3ZA/A",
            "familyType": "iPhone 13 Pro Max 256GB 金色",
            "color": "gold",
            "capacity": "256GB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$10,499.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-gold-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841927000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-gold-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841927000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLTL3ZA/A",
            "familyType": "iPhone 13 Pro 1TB 銀色",
            "color": "銀色",
            "capacity": "1TB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$13,199.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-silver-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841924000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-silver-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841924000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLTF3ZA/A",
            "familyType": "iPhone 13 Pro 512GB 石墨色",
            "color": "石墨色",
            "capacity": "512GB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$11,399.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-graphite-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629910158000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-graphite-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629910158000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLHE3ZA/A",
            "familyType": "iPhone 13 Pro Max 512GB 銀色",
            "color": "銀色",
            "capacity": "512GB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$12,299.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-silver-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841924000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-silver-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841924000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLH93ZA/A",
            "familyType": "iPhone 13 Pro Max 256GB 銀色",
            "color": "銀色",
            "capacity": "256GB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$10,499.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-silver-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841924000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-silver-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841924000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLHK3ZA/A",
            "familyType": "iPhone 13 Pro Max 1TB 金色",
            "color": "gold",
            "capacity": "1TB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$14,099.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-gold-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841927000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-gold-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841927000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLHH3ZA/A",
            "familyType": "iPhone 13 Pro Max 1TB 石墨色",
            "color": "石墨色",
            "capacity": "1TB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$14,099.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-graphite-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629910158000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-graphite-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629910158000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLH63ZA/A",
            "familyType": "iPhone 13 Pro Max 128GB 金色",
            "color": "gold",
            "capacity": "128GB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$9,699.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-gold-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841927000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-gold-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841927000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLHL3ZA/A",
            "familyType": "iPhone 13 Pro Max 1TB 天峰藍色",
            "color": "天峰藍色",
            "capacity": "1TB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$14,099.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-blue-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841926000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-blue-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841926000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLH53ZA/A",
            "familyType": "iPhone 13 Pro Max 128GB 銀色",
            "color": "銀色",
            "capacity": "128GB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$9,699.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-silver-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841924000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-silver-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841924000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLTD3ZA/A",
            "familyType": "iPhone 13 Pro 256GB 金色",
            "color": "gold",
            "capacity": "256GB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$9,699.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-gold-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841924000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-gold-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841924000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLT73ZA/A",
            "familyType": "iPhone 13 Pro 128GB 金色",
            "color": "gold",
            "capacity": "128GB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$8,799.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-gold-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841924000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-gold-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841924000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLTG3ZA/A",
            "familyType": "iPhone 13 Pro 512GB 銀色",
            "color": "銀色",
            "capacity": "512GB",
            "seoUrlToken": '6.1 吋顯示器<sup>1</sup>',
            "price": "MOP$11,399.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-silver-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841924000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-silver-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841924000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1216593290",
            "subfamily": "iPhone 13 Pro"
        },
        {
            "partNumber": "MLHJ3ZA/A",
            "familyType": "iPhone 13 Pro Max 1TB 銀色",
            "color": "銀色",
            "capacity": "1TB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$14,099.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-silver-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629841924000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-silver-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629841924000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        },
        {
            "partNumber": "MLH83ZA/A",
            "familyType": "iPhone 13 Pro Max 256GB 石墨色",
            "color": "石墨色",
            "capacity": "256GB",
            "seoUrlToken": '6.7 吋顯示器<sup>1</sup>',
            "price": "MOP$10,499.00",
            "installmentPrice": "",
            "installmentPeriod": "0",
            "paymentPrice": "",
            "paymentInstallments": "",
            "iUPPrice": "",
            "iUPInstallments": "2",
            "colorSortOrder": "",
            "swatchImage": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-graphite-select_SW_COLOR?wid=16&hei=16&fmt=jpeg&qlt=80&.v=1629910158000",
            "image": "https://store.storeimages.cdn-apple.com/4982/as-images.apple.com/is/iphone-13-pro-max-graphite-select?wid=300&hei=300&fmt=jpeg&qlt=80&.v=1629910158000",
            "contractEnabled": true,
            "unlockedEnabled": true,
            "groupID": "702344648",
            "groupName": "iPhonePro",
            "subfamilyID": "1040556462",
            "subfamily": "iPhone 13 Pro Max"
        }]`,
    },
}