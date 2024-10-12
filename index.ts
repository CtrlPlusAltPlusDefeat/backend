/* Do not change, this code is generated from Golang structs */


export interface Wrapper {
    service: string;
    action: string;
    data: string;
}
export interface ErrorResponse {
    error: string;
}
export interface SessionResponse {
    sessionId: string;
}
export interface SendChatResponse {
    text: string;
    timestamp: number;
    playerId: string;
}
export interface LoadChatResponse {
    messages: SendChatResponse[];
}
export interface JoinResponse {
    lobbyId: string;
}
export interface Settings {
    gameId: number;
    maxPlayers: number;
    teams: number;
    game: number[];
}
export interface Player {
    id: string;
    name: string;
    points: number;
    isAdmin: boolean;
    isOnline: boolean;
}
export interface Details {
    players: Player[];
    lobbyId: string;
    settings: Settings;
    inGame: boolean;
    gameSessionId: string;
}
export interface GetResponse {
    lobby: Details;
    player: Player;
}
export interface PlayerJoinResponse {
    player: Player;
}
export interface PlayerLeftResponse {
    player: Player;
}