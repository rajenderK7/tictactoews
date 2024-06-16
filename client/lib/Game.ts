const WEBSOCKET_URL = import.meta.env.VITE_WS_URL;
const NEW_GAME_ENDPOINT = WEBSOCKET_URL + import.meta.env.VITE_WS_NEW_GAME_PATH;

export interface Message {
  type: string;
  payload: string;
}

type MessageTypeAllowed =
  | "INIT GAME"
  | "START"
  | "MOVE"
  | "DRAW"
  | "GAME OVER"
  | "OK"
  | "ERROR"
  | "WAITING"
  | "DISCONNECTED";

type CallbackFn = ({ type, payload }: Message) => void;

// Game is a Singleton
class Game {
  private static instance: Game;
  private socket: WebSocket;
  private callbacks: Map<string, Function>;

  static getInstance() {
    if (!this.instance) {
      this.instance = new Game();
    }
    return this.instance;
  }

  constructor() {
    this.socket = new WebSocket(NEW_GAME_ENDPOINT);
    this.callbacks = new Map<string, CallbackFn>();
    this.attachHandlers();
  }

  private attachHandlers = () => {
    // @ts-ignore
    this.socket.onopen = (ev) => {
      console.log("Connected to the WebSocket server");
    };

    this.socket.onmessage = (ev) => {
      this.socketNewMessage(JSON.parse(ev.data) as Message);
    };

    // @ts-ignore
    this.socket.onclose = (ev) => {
      console.log("Disconnected from the WebSocket server");
    };

    this.socket.onerror = (ev) => {
      console.error("WebSocket error:", ev);
    };
  };

  public addCallback = (type: string, fn: Function) => {
    this.callbacks.set(type, fn);
  };

  public removeAllCallbacks = () => {
    this.callbacks.clear();
  };

  private socketNewMessage = (data: Message) => {
    if (!this.callbacks.has(data.type)) {
      console.log("Undefined behaviour: ", data.type);
      return;
    }

    this.callbacks.get(data.type)!!(data);
  };

  public sendMessage = (type: MessageTypeAllowed, payload: string) => {
    if (!type || !payload) return;
    this.socket.send(
      JSON.stringify({
        type,
        payload,
      })
    );
  };
}

export const GameInstance = Game.getInstance();
