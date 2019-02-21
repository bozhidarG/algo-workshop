class RealApi {
	constructor() {
		this.users = ["user1", "user2", "user3"];
	}

	getUser(userId) {
		// can have some requests towards an real api and do what it needs
		return this.users[userId] + " real";
	}
}

module.exports = RealApi;