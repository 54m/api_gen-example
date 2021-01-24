// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: 1.7.1

export type User = {
	age: number;
	gender: number;
	id: string;
	name: string;
}
export type FailedReason = {
	message: string;
}
export type PostAgeIncrementRequest = {
	userID: string;
}
export type PostAgeIncrementResponse = {
	messages: FailedReason[] | null;
	payload?: User;
	status: number;
}
