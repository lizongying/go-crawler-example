package levi_spider

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-json/gjson"
)

type Spider struct {
	pkg.Spider
	logger             pkg.Logger
	collectionCategory string
	collectionDetail   string
}

func (s *Spider) ParseList(ctx pkg.Context, response pkg.Response) (err error) {
	js := response.MustJson()
	for _, v := range js.ManySelector(`data.categories.products`) {
		data := DataDetail{
			Id:       v.One(`code`).String(),
			Title:    v.One(`name`).String(),
			Subtitle: v.One(`department`).String(),
		}
		s.MustYieldItem(ctx, items.NewItemMongo(s.collectionDetail, true).
			SetUniqueKey(data.Id).
			SetId(data.Id).
			SetData(&data))
	}
	if js.One(`data.categories.pagination.currentPage`).String() != js.One(`data.categories.pagination.totalPages`).String() {
		jsStr := map[string]any{
			"operationName": "categories",
			"query":         "query categories($query: String!, $sort: String, $currentPage: Int, $pageSize: Int, $categoryId: String!, $preOrder: Boolean, $abTestSectionedPLP: Boolean, $enableGrouping: Boolean) {\n  categories(\n    query: $query\n    sort: $sort\n    currentPage: $currentPage\n    pageSize: $pageSize\n    categoryId: $categoryId\n    preOrder: $preOrder\n    abTestSectionedPLP: $abTestSectionedPLP\n    enableGrouping: $enableGrouping\n  ) {\n    description\n    breadcrumbs {\n      facetCode\n      facetName\n      facetValueName\n      removeQuery {\n        query {\n          value\n        }\n        url\n      }\n    }\n    categoryCode\n    categoryHierarchy {\n      code\n      count\n      selected\n      childSelected\n      children\n      depth\n      leaf\n      parentSelected\n    }\n    categoryName\n    emailSignUpGateEnabled\n    registrationGateEnabled\n    currentQuery {\n      url\n    }\n    facets {\n      category\n      code\n      name\n      nofollow\n      priority\n      visible\n      topValues {\n        count\n        name\n        nofollow\n        selected\n        query {\n          query {\n            value\n          }\n          url\n        }\n      }\n      values {\n        count\n        name\n        nofollow\n        selected\n        query {\n          query {\n            value\n          }\n          url\n        }\n      }\n    }\n    faq\n    freeTextSearch\n    noProductsRedirectMsg\n    lscoBreadcrumbs {\n      name\n      url\n      linkClass\n    }\n    pagination {\n      currentPage\n      totalPages\n      totalResults\n    }\n    products {\n      channels\n      code\n      backOrder\n      name\n      url\n      price {\n        code\n        currencyIso\n        formattedValue\n        hardPrice\n        hardPriceFormattedValue\n        regularPrice\n        regularPriceFormattedValue\n        softPrice\n        softPriceFormattedValue\n        value\n      }\n      priceRange {\n        maxPrice {\n          formattedValue\n          value\n          regularPrice\n          softPrice\n          hardPrice\n        }\n        minPrice {\n          formattedValue\n          value\n          regularPrice\n          softPrice\n          hardPrice\n        }\n      }\n      priceRangeFrom {\n        maxPrice {\n          formattedValue\n          value\n          regularPrice\n          softPrice\n          hardPrice\n        }\n        minPrice {\n          formattedValue\n          value\n          regularPrice\n          softPrice\n          hardPrice\n        }\n      }\n      baseProduct\n      soldIndividually\n      comingSoon\n      averageOverallRatings\n      noOfRatings\n      soldOutForever\n      sustainability\n      findInStoreEligible\n      customizable\n      flxCustomization\n      availableForPickup\n      department\n      productSizeCoverage\n      pdpGroupId\n      preOrder\n      preOrderShipDate\n      returnable\n      variantOptions {\n        code\n        comingSoon\n        preOrder\n        backOrder\n        customizable\n        findInStoreEligible\n        flxCustomization\n        merchantBadge\n        promotionalBadge\n        sustainability\n        name\n        swatchUrl\n        swatchAltText\n        galleryList {\n          galleryImage {\n            altText\n            format\n            galleryIndex\n            imageType\n            url\n          }\n        }\n        priceData {\n          hardPrice\n          hardPriceFormattedValue\n          regularPrice\n          regularPriceFormattedValue\n          softPrice\n          softPriceFormattedValue\n          value\n          currencyIso\n        }\n        soldIndividually\n        soldOutForever\n        url\n      }\n      lscoBreadcrumbs {\n        categoryCode\n        name\n        url\n      }\n      swatchUrl\n      swatchAltText\n      galleryList {\n        galleryImage {\n          altText\n          format\n          galleryIndex\n          imageType\n          url\n        }\n      }\n      merchantBadge\n      promotionalBadge\n      errors {\n        component\n        name\n        time_thrown\n        message\n      }\n    }\n    seoMetaData {\n      canonicalUrl\n      metaDescription\n      metaH1\n      metaTitle\n      robots\n    }\n    sorts {\n      code\n      name\n      selected\n    }\n    spellingSuggestion {\n      query\n      suggestion\n    }\n    subSections {\n      description\n      categoryName\n      currentQuery {\n        url\n      }\n      pagination {\n        currentPage\n        totalPages\n        totalResults\n      }\n      products {\n        channels\n        code\n        backOrder\n        name\n        url\n        price {\n          code\n          currencyIso\n          formattedValue\n          hardPrice\n          hardPriceFormattedValue\n          regularPrice\n          regularPriceFormattedValue\n          softPrice\n          softPriceFormattedValue\n          value\n        }\n        priceRange {\n          maxPrice {\n            formattedValue\n            value\n            regularPrice\n            softPrice\n            hardPrice\n          }\n          minPrice {\n            formattedValue\n            value\n            regularPrice\n            softPrice\n            hardPrice\n          }\n        }\n        priceRangeFrom {\n          maxPrice {\n            formattedValue\n            value\n            regularPrice\n            softPrice\n            hardPrice\n          }\n          minPrice {\n            formattedValue\n            value\n            regularPrice\n            softPrice\n            hardPrice\n          }\n        }\n        baseProduct\n        soldIndividually\n        comingSoon\n        averageOverallRatings\n        noOfRatings\n        soldOutForever\n        sustainability\n        findInStoreEligible\n        customizable\n        flxCustomization\n        availableForPickup\n        department\n        pdpGroupId\n        preOrder\n        preOrderShipDate\n        returnable\n        variantOptions {\n          code\n          comingSoon\n          preOrder\n          backOrder\n          customizable\n          findInStoreEligible\n          flxCustomization\n          merchantBadge\n          promotionalBadge\n          sustainability\n          name\n          swatchUrl\n          swatchAltText\n          galleryList {\n            galleryImage {\n              altText\n              format\n              galleryIndex\n              imageType\n              url\n            }\n          }\n          priceData {\n            hardPrice\n            hardPriceFormattedValue\n            regularPrice\n            regularPriceFormattedValue\n            softPrice\n            softPriceFormattedValue\n            value\n            currencyIso\n          }\n          soldIndividually\n          soldOutForever\n          url\n        }\n        lscoBreadcrumbs {\n          categoryCode\n          name\n          url\n        }\n        swatchUrl\n        swatchAltText\n        galleryList {\n          galleryImage {\n            altText\n            format\n            galleryIndex\n            imageType\n            url\n          }\n        }\n        merchantBadge\n        promotionalBadge\n        errors {\n          component\n          name\n          time_thrown\n          message\n        }\n      }\n      seoMetaData {\n        canonicalUrl\n        metaDescription\n        metaH1\n        metaTitle\n        robots\n      }\n    }\n  }\n}\n",
			"variables": map[string]any{
				"query":              ":relevance",
				"country":            "US",
				"locale":             "en_US",
				"currentPage":        js.One(`data.categories.pagination.currentPage`).Int() + 1,
				"pageSize":           js.One(`data.categories.pagination.pageSize`).Int(),
				"sort":               "relevance",
				"categoryId":         "levi_clothing",
				"preOrder":           true,
				"abTestSectionedPLP": false,
				"enableGrouping":     true,
			},
		}
		s.MustYieldRequest(ctx, request.NewRequest().
			SetUrl("https://www.levi.com/nextgen-webhooks/").
			SetQuery("operationName", "categories").
			SetQuery("locale", "US-en_US").
			SetBodyJson(jsStr).
			SetHeader("x-brand", "levi").
			SetHeader("x-country", "US").
			SetHeader("x-locale", "en_US").
			SetHeader("x-operationname", "categories").
			SetProxy("http://127.0.0.1:7890").
			SetCallBack(s.ParseList))
	}
	return
}

func (s *Spider) ParseIndex(ctx pkg.Context, response pkg.Response) (err error) {
	jsonStr := response.MustReOne(`window\.__LSCO_INITIAL_STATE__ = Object\.freeze\((.+?)\);`).String()
	js, err := gjson.NewSelectorFromStr(jsonStr)
	if err != nil {
		return err
	}

	for _, v := range js.ManySelector(`ssrViewStoreProductList.products`) {
		data := DataDetail{
			Id:       v.One(`code`).String(),
			Title:    v.One(`name`).String(),
			Subtitle: v.One(`department`).String(),
		}
		s.MustYieldItem(ctx, items.NewItemMongo(s.collectionDetail, true).
			SetUniqueKey(data.Id).
			SetId(data.Id).
			SetData(&data))
	}

	if js.One(`ssrViewStoreProductList.pagination.currentPage`).String() != js.One(`ssrViewStoreProductList.pagination.totalPages`).String() {
		jsStr := map[string]any{
			"operationName": "categories",
			"query":         "query categories($query: String!, $sort: String, $currentPage: Int, $pageSize: Int, $categoryId: String!, $preOrder: Boolean, $abTestSectionedPLP: Boolean, $enableGrouping: Boolean) {\n  categories(\n    query: $query\n    sort: $sort\n    currentPage: $currentPage\n    pageSize: $pageSize\n    categoryId: $categoryId\n    preOrder: $preOrder\n    abTestSectionedPLP: $abTestSectionedPLP\n    enableGrouping: $enableGrouping\n  ) {\n    description\n    breadcrumbs {\n      facetCode\n      facetName\n      facetValueName\n      removeQuery {\n        query {\n          value\n        }\n        url\n      }\n    }\n    categoryCode\n    categoryHierarchy {\n      code\n      count\n      selected\n      childSelected\n      children\n      depth\n      leaf\n      parentSelected\n    }\n    categoryName\n    emailSignUpGateEnabled\n    registrationGateEnabled\n    currentQuery {\n      url\n    }\n    facets {\n      category\n      code\n      name\n      nofollow\n      priority\n      visible\n      topValues {\n        count\n        name\n        nofollow\n        selected\n        query {\n          query {\n            value\n          }\n          url\n        }\n      }\n      values {\n        count\n        name\n        nofollow\n        selected\n        query {\n          query {\n            value\n          }\n          url\n        }\n      }\n    }\n    faq\n    freeTextSearch\n    noProductsRedirectMsg\n    lscoBreadcrumbs {\n      name\n      url\n      linkClass\n    }\n    pagination {\n      currentPage\n      totalPages\n      totalResults\n    }\n    products {\n      channels\n      code\n      backOrder\n      name\n      url\n      price {\n        code\n        currencyIso\n        formattedValue\n        hardPrice\n        hardPriceFormattedValue\n        regularPrice\n        regularPriceFormattedValue\n        softPrice\n        softPriceFormattedValue\n        value\n      }\n      priceRange {\n        maxPrice {\n          formattedValue\n          value\n          regularPrice\n          softPrice\n          hardPrice\n        }\n        minPrice {\n          formattedValue\n          value\n          regularPrice\n          softPrice\n          hardPrice\n        }\n      }\n      priceRangeFrom {\n        maxPrice {\n          formattedValue\n          value\n          regularPrice\n          softPrice\n          hardPrice\n        }\n        minPrice {\n          formattedValue\n          value\n          regularPrice\n          softPrice\n          hardPrice\n        }\n      }\n      baseProduct\n      soldIndividually\n      comingSoon\n      averageOverallRatings\n      noOfRatings\n      soldOutForever\n      sustainability\n      findInStoreEligible\n      customizable\n      flxCustomization\n      availableForPickup\n      department\n      productSizeCoverage\n      pdpGroupId\n      preOrder\n      preOrderShipDate\n      returnable\n      variantOptions {\n        code\n        comingSoon\n        preOrder\n        backOrder\n        customizable\n        findInStoreEligible\n        flxCustomization\n        merchantBadge\n        promotionalBadge\n        sustainability\n        name\n        swatchUrl\n        swatchAltText\n        galleryList {\n          galleryImage {\n            altText\n            format\n            galleryIndex\n            imageType\n            url\n          }\n        }\n        priceData {\n          hardPrice\n          hardPriceFormattedValue\n          regularPrice\n          regularPriceFormattedValue\n          softPrice\n          softPriceFormattedValue\n          value\n          currencyIso\n        }\n        soldIndividually\n        soldOutForever\n        url\n      }\n      lscoBreadcrumbs {\n        categoryCode\n        name\n        url\n      }\n      swatchUrl\n      swatchAltText\n      galleryList {\n        galleryImage {\n          altText\n          format\n          galleryIndex\n          imageType\n          url\n        }\n      }\n      merchantBadge\n      promotionalBadge\n      errors {\n        component\n        name\n        time_thrown\n        message\n      }\n    }\n    seoMetaData {\n      canonicalUrl\n      metaDescription\n      metaH1\n      metaTitle\n      robots\n    }\n    sorts {\n      code\n      name\n      selected\n    }\n    spellingSuggestion {\n      query\n      suggestion\n    }\n    subSections {\n      description\n      categoryName\n      currentQuery {\n        url\n      }\n      pagination {\n        currentPage\n        totalPages\n        totalResults\n      }\n      products {\n        channels\n        code\n        backOrder\n        name\n        url\n        price {\n          code\n          currencyIso\n          formattedValue\n          hardPrice\n          hardPriceFormattedValue\n          regularPrice\n          regularPriceFormattedValue\n          softPrice\n          softPriceFormattedValue\n          value\n        }\n        priceRange {\n          maxPrice {\n            formattedValue\n            value\n            regularPrice\n            softPrice\n            hardPrice\n          }\n          minPrice {\n            formattedValue\n            value\n            regularPrice\n            softPrice\n            hardPrice\n          }\n        }\n        priceRangeFrom {\n          maxPrice {\n            formattedValue\n            value\n            regularPrice\n            softPrice\n            hardPrice\n          }\n          minPrice {\n            formattedValue\n            value\n            regularPrice\n            softPrice\n            hardPrice\n          }\n        }\n        baseProduct\n        soldIndividually\n        comingSoon\n        averageOverallRatings\n        noOfRatings\n        soldOutForever\n        sustainability\n        findInStoreEligible\n        customizable\n        flxCustomization\n        availableForPickup\n        department\n        pdpGroupId\n        preOrder\n        preOrderShipDate\n        returnable\n        variantOptions {\n          code\n          comingSoon\n          preOrder\n          backOrder\n          customizable\n          findInStoreEligible\n          flxCustomization\n          merchantBadge\n          promotionalBadge\n          sustainability\n          name\n          swatchUrl\n          swatchAltText\n          galleryList {\n            galleryImage {\n              altText\n              format\n              galleryIndex\n              imageType\n              url\n            }\n          }\n          priceData {\n            hardPrice\n            hardPriceFormattedValue\n            regularPrice\n            regularPriceFormattedValue\n            softPrice\n            softPriceFormattedValue\n            value\n            currencyIso\n          }\n          soldIndividually\n          soldOutForever\n          url\n        }\n        lscoBreadcrumbs {\n          categoryCode\n          name\n          url\n        }\n        swatchUrl\n        swatchAltText\n        galleryList {\n          galleryImage {\n            altText\n            format\n            galleryIndex\n            imageType\n            url\n          }\n        }\n        merchantBadge\n        promotionalBadge\n        errors {\n          component\n          name\n          time_thrown\n          message\n        }\n      }\n      seoMetaData {\n        canonicalUrl\n        metaDescription\n        metaH1\n        metaTitle\n        robots\n      }\n    }\n  }\n}\n",
			"variables": map[string]any{
				"query":              ":relevance",
				"country":            "US",
				"locale":             "en_US",
				"currentPage":        js.One(`ssrViewStoreProductList.pagination.currentPage`).Int() + 1,
				"pageSize":           js.One(`ssrViewStoreProductList.pagination.pageSize`).Int(),
				"sort":               "relevance",
				"categoryId":         "levi_clothing",
				"preOrder":           true,
				"abTestSectionedPLP": false,
				"enableGrouping":     true,
			},
		}
		s.MustYieldRequest(ctx, request.NewRequest().
			SetUrl("https://www.levi.com/nextgen-webhooks/").
			SetQuery("operationName", "categories").
			SetQuery("locale", "US-en_US").
			SetBodyJson(jsStr).
			SetHeader("x-brand", "levi").
			SetHeader("x-country", "US").
			SetHeader("x-locale", "en_US").
			SetHeader("x-operationname", "categories").
			SetProxy("http://127.0.0.1:7890").
			SetCallBack(s.ParseList))
	}
	return
}

// TestList go run cmd/levi_spider/*.go -c dev.yml -n levi -f TestList -m once
func (s *Spider) TestList(ctx pkg.Context, _ string) (err error) {
	js := map[string]any{
		"operationName": "categories",
		"query":         "query categories($query: String!, $sort: String, $currentPage: Int, $pageSize: Int, $categoryId: String!, $preOrder: Boolean, $abTestSectionedPLP: Boolean, $enableGrouping: Boolean) {\n  categories(\n    query: $query\n    sort: $sort\n    currentPage: $currentPage\n    pageSize: $pageSize\n    categoryId: $categoryId\n    preOrder: $preOrder\n    abTestSectionedPLP: $abTestSectionedPLP\n    enableGrouping: $enableGrouping\n  ) {\n    description\n    breadcrumbs {\n      facetCode\n      facetName\n      facetValueName\n      removeQuery {\n        query {\n          value\n        }\n        url\n      }\n    }\n    categoryCode\n    categoryHierarchy {\n      code\n      count\n      selected\n      childSelected\n      children\n      depth\n      leaf\n      parentSelected\n    }\n    categoryName\n    emailSignUpGateEnabled\n    registrationGateEnabled\n    currentQuery {\n      url\n    }\n    facets {\n      category\n      code\n      name\n      nofollow\n      priority\n      visible\n      topValues {\n        count\n        name\n        nofollow\n        selected\n        query {\n          query {\n            value\n          }\n          url\n        }\n      }\n      values {\n        count\n        name\n        nofollow\n        selected\n        query {\n          query {\n            value\n          }\n          url\n        }\n      }\n    }\n    faq\n    freeTextSearch\n    noProductsRedirectMsg\n    lscoBreadcrumbs {\n      name\n      url\n      linkClass\n    }\n    pagination {\n      currentPage\n      totalPages\n      totalResults\n    }\n    products {\n      channels\n      code\n      backOrder\n      name\n      url\n      price {\n        code\n        currencyIso\n        formattedValue\n        hardPrice\n        hardPriceFormattedValue\n        regularPrice\n        regularPriceFormattedValue\n        softPrice\n        softPriceFormattedValue\n        value\n      }\n      priceRange {\n        maxPrice {\n          formattedValue\n          value\n          regularPrice\n          softPrice\n          hardPrice\n        }\n        minPrice {\n          formattedValue\n          value\n          regularPrice\n          softPrice\n          hardPrice\n        }\n      }\n      priceRangeFrom {\n        maxPrice {\n          formattedValue\n          value\n          regularPrice\n          softPrice\n          hardPrice\n        }\n        minPrice {\n          formattedValue\n          value\n          regularPrice\n          softPrice\n          hardPrice\n        }\n      }\n      baseProduct\n      soldIndividually\n      comingSoon\n      averageOverallRatings\n      noOfRatings\n      soldOutForever\n      sustainability\n      findInStoreEligible\n      customizable\n      flxCustomization\n      availableForPickup\n      department\n      productSizeCoverage\n      pdpGroupId\n      preOrder\n      preOrderShipDate\n      returnable\n      variantOptions {\n        code\n        comingSoon\n        preOrder\n        backOrder\n        customizable\n        findInStoreEligible\n        flxCustomization\n        merchantBadge\n        promotionalBadge\n        sustainability\n        name\n        swatchUrl\n        swatchAltText\n        galleryList {\n          galleryImage {\n            altText\n            format\n            galleryIndex\n            imageType\n            url\n          }\n        }\n        priceData {\n          hardPrice\n          hardPriceFormattedValue\n          regularPrice\n          regularPriceFormattedValue\n          softPrice\n          softPriceFormattedValue\n          value\n          currencyIso\n        }\n        soldIndividually\n        soldOutForever\n        url\n      }\n      lscoBreadcrumbs {\n        categoryCode\n        name\n        url\n      }\n      swatchUrl\n      swatchAltText\n      galleryList {\n        galleryImage {\n          altText\n          format\n          galleryIndex\n          imageType\n          url\n        }\n      }\n      merchantBadge\n      promotionalBadge\n      errors {\n        component\n        name\n        time_thrown\n        message\n      }\n    }\n    seoMetaData {\n      canonicalUrl\n      metaDescription\n      metaH1\n      metaTitle\n      robots\n    }\n    sorts {\n      code\n      name\n      selected\n    }\n    spellingSuggestion {\n      query\n      suggestion\n    }\n    subSections {\n      description\n      categoryName\n      currentQuery {\n        url\n      }\n      pagination {\n        currentPage\n        totalPages\n        totalResults\n      }\n      products {\n        channels\n        code\n        backOrder\n        name\n        url\n        price {\n          code\n          currencyIso\n          formattedValue\n          hardPrice\n          hardPriceFormattedValue\n          regularPrice\n          regularPriceFormattedValue\n          softPrice\n          softPriceFormattedValue\n          value\n        }\n        priceRange {\n          maxPrice {\n            formattedValue\n            value\n            regularPrice\n            softPrice\n            hardPrice\n          }\n          minPrice {\n            formattedValue\n            value\n            regularPrice\n            softPrice\n            hardPrice\n          }\n        }\n        priceRangeFrom {\n          maxPrice {\n            formattedValue\n            value\n            regularPrice\n            softPrice\n            hardPrice\n          }\n          minPrice {\n            formattedValue\n            value\n            regularPrice\n            softPrice\n            hardPrice\n          }\n        }\n        baseProduct\n        soldIndividually\n        comingSoon\n        averageOverallRatings\n        noOfRatings\n        soldOutForever\n        sustainability\n        findInStoreEligible\n        customizable\n        flxCustomization\n        availableForPickup\n        department\n        pdpGroupId\n        preOrder\n        preOrderShipDate\n        returnable\n        variantOptions {\n          code\n          comingSoon\n          preOrder\n          backOrder\n          customizable\n          findInStoreEligible\n          flxCustomization\n          merchantBadge\n          promotionalBadge\n          sustainability\n          name\n          swatchUrl\n          swatchAltText\n          galleryList {\n            galleryImage {\n              altText\n              format\n              galleryIndex\n              imageType\n              url\n            }\n          }\n          priceData {\n            hardPrice\n            hardPriceFormattedValue\n            regularPrice\n            regularPriceFormattedValue\n            softPrice\n            softPriceFormattedValue\n            value\n            currencyIso\n          }\n          soldIndividually\n          soldOutForever\n          url\n        }\n        lscoBreadcrumbs {\n          categoryCode\n          name\n          url\n        }\n        swatchUrl\n        swatchAltText\n        galleryList {\n          galleryImage {\n            altText\n            format\n            galleryIndex\n            imageType\n            url\n          }\n        }\n        merchantBadge\n        promotionalBadge\n        errors {\n          component\n          name\n          time_thrown\n          message\n        }\n      }\n      seoMetaData {\n        canonicalUrl\n        metaDescription\n        metaH1\n        metaTitle\n        robots\n      }\n    }\n  }\n}\n",
		"variables": map[string]any{
			"query":              ":relevance",
			"country":            "US",
			"locale":             "en_US",
			"currentPage":        3,
			"pageSize":           36,
			"sort":               "relevance",
			"categoryId":         "levi_clothing",
			"preOrder":           true,
			"abTestSectionedPLP": false,
			"enableGrouping":     true,
		},
	}
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl("https://www.levi.com/nextgen-webhooks/").
		SetQuery("operationName", "categories").
		SetQuery("locale", "US-en_US").
		SetBodyJson(js).
		SetHeader("x-brand", "levi").
		SetHeader("x-country", "US").
		SetHeader("x-locale", "en_US").
		SetHeader("x-operationname", "categories").
		SetProxy("http://127.0.0.1:7890").
		SetCallBack(s.ParseList))
	return
}

// TestIndex go run cmd/levi_spider/*.go -c dev.yml -n levi -f TestIndex -m once
func (s *Spider) TestIndex(ctx pkg.Context, _ string) (err error) {
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl("https://www.levi.com/US/en_US/clothing/c/levi_clothing").
		SetProxy("http://127.0.0.1:7890").
		SetCallBack(s.ParseIndex))
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider:             baseSpider,
		logger:             baseSpider.GetLogger(),
		collectionCategory: "levi_category",
		collectionDetail:   "levi_detail",
	}
	spider.WithOptions(
		pkg.WithName("levi"),
		pkg.WithMongoPipeline(),
		pkg.WithRetryMaxTimes(0),
	)
	return
}
