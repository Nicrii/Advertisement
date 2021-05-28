package domain

//func TestGetAdNoAdFound(t *testing.T) {
//	ad, err := AdDao.Get(0)
//	assert.Nil(t, ad, "we were not expecting a ad with id = 0")
//	assert.NotNil(t, err, "we were not expecting an error when ad id is 0")
//	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
//	assert.EqualValues(t, "not_found", err.Code)
//	assert.EqualValues(t, "ad 0 not found", err.Message)
//}
//
//func TestGetAdNoError(t *testing.T) {
//	ad, err := AdDao.Get(123)
//	assert.Nil(t, err)
//	assert.NotNil(t, ad)
//	assert.EqualValues(t, ad.Id, uint64(123))
//	assert.EqualValues(t, ad.Name, "TestAdvertisement")
//	assert.EqualValues(t, ad.Description, "This is test date")
//	assert.EqualValues(t, ad.ImagesURLs, []string{"link1", "link2"})
//	assert.EqualValues(t, ad.Price, 100500)
//	assert.EqualValues(t, ad.Date, date.Date{Year: 2021})
//}
