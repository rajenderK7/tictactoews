import { useEffect, useState } from "react";
import { GameInstance, Message } from "../lib/Game";
import { Mark, MessageType } from "../lib/constants";
import Cell from "./components/Cell";

const Game = () => {
  const [playerMark, setPlayerMark] = useState<string>(Mark.X);
  const [isGameOver, setIsGameOver] = useState<boolean>(false);
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
  const handleWaiting = ({ payload }: Message) => {
    updateMsg(payload);
  };

  const handleStart = ({ payload }: Message) => {
    setPlayerMark(payload);
    clearMsg();
    setIsGameOver(false);
  };

  const handleMove = (rowIdx: number, colIdx: number) => {
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
  };

  const handleError = ({ payload }: Message) => {
    updateMsg(payload);
  };

  const handleGameOver = ({ payload }: Message) => {
    updateMsg(payload);
    setIsGameOver(true);
  };

  const handleRefresh = () => {
    window.location.reload();
  };

  useEffect(() => {
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
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <h1 className="text-4xl font-bold mb-2 text-pink-700">Tic Tac Toe</h1>
      <div className="text-xl font-semibold text-gray-800 mb-4">
        Your Mark:{" "}
        <span className="text-green-600 font-bold">{playerMark}</span>
      </div>
      {msg.length > 0 && (
        <p className="text-white bg-black px-4 py-1 rounded mb-6 shadow-md">
          {msg}
        </p>
      )}
      <div className="flex flex-col items-center border-2 border-black shadow-xl bg-white">
        {board.map((row, rowIdx) => (
          <div key={rowIdx} className="flex">
            {row.map((mark, colIdx) => (
              <Cell
                key={`${rowIdx}-${colIdx}`}
                onClick={handleMove}
                mark={mark}
                rowIdx={rowIdx}
                colIdx={colIdx}
              />
            ))}
          </div>
        ))}
      </div>
      {isGameOver && (
        <button
          onClick={handleRefresh}
          className="mt-4 px-3 py-2 bg-gray-600 text-white rounded shadow-md hover:bg-gray-700 transition-colors"
        >
          Refresh ↻
        </button>
      )}
    </div>
  );
};

export default Game;
