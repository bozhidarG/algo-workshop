class MockApi {
	constructor() {
		this.users = ["user1", "user2", "user3"];
	}

	getUser(userId) {
		// can have some test logic, when used for tests
		return this.users[userId] + " mock";
	}
}

module.exports = MockApi;