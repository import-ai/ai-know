interface NoteEditorProps {
  content: string,
  onUpdate: (newContent: string) => void
}

const NoteEditor: React.FC<NoteEditorProps> = ({ content, onUpdate }) => {
  return (
    <div className="flex-grow p-4">
      <textarea
        className="w-full h-full p-4 border border-gray-300 resize-none"
        placeholder="Write your note here..."
        value={content}
        onChange={(e) => onUpdate(e.target.value)}
      >
      </textarea>
    </div>
  );
}

export default NoteEditor