import { useEffect, useState } from "react";
import { GameInstance, Message } from "../lib/Game";
import { Mark, MessageType } from "../lib/constants";
import Cell from "./components/Cell";

const Game = () => {
  const [playerMark, setPlayerMark] = useState<string>(Mark.X);
  // const [opponentMark, setOpponentMark] = useState<string>(Mark.O);
  const [playerTurn, setPlayerTurn] = useState<boolean>(true);
  const [msg, setMsg] = useState<string>("");
  const [board, setBoard] = useState<string[][]>([
    ["", "", ""],
    ["", "", ""],
    ["", "", ""],
  ]);

  const updateBoard = (mark: string, rowIdx: number, colIdx: number) => {
    setBoard((prev) => {
      const newState = [...prev];
      newState[rowIdx][colIdx] = mark;
      return newState;
    });
  };

  const updateMsg = (newMsg: string) => {
    if (newMsg && newMsg.length > 0) {
      setMsg(newMsg);
    }
  };

  const clearMsg = () => {
    setMsg("");
  };

  const switchTurn = () => {
    setPlayerTurn((prev) => !prev);
  };

  const handleWaiting = ({ payload }: Message) => {
    updateMsg(payload);
  };

  const handleStart = ({ payload }: Message) => {
    setPlayerMark(payload);
    if (payload === Mark.X) {
      setPlayerTurn(true);
    } else {
      setPlayerTurn(false);
    }
  };

  const handleMove = (rowIdx: number, colIdx: number) => {
    // if (!playerTurn) {
    //   console.log("Inside");
    //   return;
    // }
    switchTurn();
    GameInstance.sendMessage(MessageType.MOVE, `${rowIdx} ${colIdx}`);
  };

  const handleRecieveMove = ({ payload }: Message) => {
    // payload will be of the form "mark row col". Eg: "X 0 2" or "O 2 1"
    const move = payload.split(" ");
    const mark = move[0];
    const rowIdx = parseInt(move[1]);
    const colIdx = parseInt(move[2]);

    clearMsg();
    updateBoard(mark, rowIdx, colIdx);
    switchTurn();
  };

  const handleError = ({ payload }: Message) => {
    updateMsg(payload);
  };

  const handleGameOver = ({ payload }: Message) => {
    updateMsg(payload);
    GameInstance.sendMessage("MOVE", "0 0");
  };

  const handleRefresh = () => {
    GameInstance.sendMessage(MessageType.INIT_GAME, "");
  };

  useEffect(() => {
    // GameInstance.addCallback(MessageType.INIT_GAME, handleInitGame);
    GameInstance.addCallback(MessageType.START, handleStart);
    GameInstance.addCallback(MessageType.WAITING, handleWaiting);
    GameInstance.addCallback(MessageType.MOVE, handleRecieveMove);
    GameInstance.addCallback(MessageType.GAME_OVER, handleGameOver);
    GameInstance.addCallback(MessageType.ERROR, handleError);

    return () => {
      GameInstance.removeAllCallbacks();
    };
  }, []);

  return (
    <div className="text-xl font-bold text-center">
      <h1>{playerMark}</h1>
      {msg.length > 0 && <p>{msg}</p>}
      {/* <button onClick={handleRefresh}>Refresh</button> */}
      <div className="flex flex-col items-center justify-center min-h-screen">
        <h1 className="text-4xl font-bold mb-8">Tic Tac Toe</h1>
        <div className="flex flex-col justify-center items-center border-2 border-black shadow-xl">
          {board.map((row, rowIdx) => {
            return (
              <div key={rowIdx} className="flex justify-evenly w-full">
                {row.map((mark, colIdx) => {
                  return (
                    <Cell
                      key={`${rowIdx}-${colIdx}`}
                      onClick={handleMove}
                      mark={mark}
                      rowIdx={rowIdx}
                      colIdx={colIdx}
                    />
                  );
                })}
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default Game;
