export interface CellProps {
  mark: string;
  rowIdx: number;
  colIdx: number;
  onClick: (rowIdx: number, colIdx: number) => void;
}

const Cell = ({ mark, rowIdx, colIdx, onClick }: CellProps) => {
  return (
    <div
      onClick={() => onClick(rowIdx, colIdx)}
      className="flex justify-center items-center h-24 w-24 hover:cursor-pointer bg-red-200 text-center text-3xl font-bold hover:bg-red-300 border-2 border-black text-pink-600"
    >
      {mark}
    </div>
  );
};

export default Cell;
