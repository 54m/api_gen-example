// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: 1.6.1

export type User = {
	age: number;
	gender: number;
	name: string;
}
export type FailedReason = {
	message: string;
}
export type DeleteRequest = {
	userID: string;
}
export type DeleteResponse = {
	messages: FailedReason[] | null;
	status: number;
}
export type GetRequest = {
	userID: string;
}
export type GetResponse = {
	messages: FailedReason[] | null;
	payload?: User;
	status: number;
}
export type GetSearchRequest = {
	age: number;
	gender: number;
	name: string;
}
export type GetSearchResponse = {
	messages: FailedReason[] | null;
	payload?: (User | null)[];
	status: number;
}
export type PatchRequest = {
	age?: number;
	gender?: number;
	name?: string;
	userID: string;
}
export type PatchResponse = {
	messages: FailedReason[] | null;
	payload?: User;
	status: number;
}
