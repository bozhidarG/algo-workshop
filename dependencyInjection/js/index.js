const RealApi = require("./api/realApi.js");
const MockApi = require("./api/mockApi.js");

class ApiConsumer {
	constructor(api) {
		this.api = api;
	}

	getUser(userId) {
		return this.api.getUser(userId);
	}
}

function getRealConsumer() {
	realApi = new RealApi();
	return new ApiConsumer(realApi);
}

function getMockConsumer() {
	mockApi = new MockApi();
	return new ApiConsumer(mockApi);
}

consumer = getMockConsumer();
console.log(consumer.getUser(1));

consumerReal = getRealConsumer();
console.log(consumerReal.getUser(1));
