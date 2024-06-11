import React from "react";

interface Note {
  id: number,
  content: string,
}

interface NoteListProps {
  notes: Note[],
  onSelect: (id: number) => void
}

const NoteList: React.FC<NoteListProps> = ({ notes, onSelect }) => {
  const items = notes.map((note: Note) => {
    return <li
      className="p-4 bg-white border border-gray-300 cursor-pointer"
      onClick={() => onSelect(note.id)}
      key={note.id}
    >
      {note.content.slice(0, 15)}
    </li>
  })
  return (
    <div className="w-1/4 bg-gray-100 p-4 shadow-lg overflow-y-auto">
      <ul className="space-y-4">
        <button
          className="w-full py-2 px-4 bg-green-500 text-white font-semibold rounded-lg shadow-md hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-green-400 focus:ring-opacity-75 mb-4"
          onClick={() => onSelect(-1)}
        >
          Create New Note
        </button>
        {items}
      </ul>
    </div>
  );
}

export default NoteList