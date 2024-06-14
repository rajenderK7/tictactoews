export const Mark = {
  X: "X",
  O: "O",
} as const;

export const MessageType = {
  INIT_GAME: "INIT GAME",
  START: "START",
  MOVE: "MOVE",
  DRAW: "DRAW",
  GAME_OVER: "GAME OVER",
  OK: "OK",
  ERROR: "ERROR",
  WAITING: "WAITING",
  DISCONNECTED: "DISCONNECTED",
} as const;
