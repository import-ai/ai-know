import Image from "next/image";

export default function Home() {
  return (
    <main className="flex h-screen">
      <div id="notes-list" className="w-1/4 bg-gray-100 p-4 shadow-lg overflow-y-auto">
        <h2 className="text-center text-2xl font-semibold mb-4">Notes</h2>
        <ul className="space-y-4">
          <li className="p-4 bg-white border border-gray-300 cursor-pointer">Note 1</li>
          <li className="p-4 bg-white border border-gray-300 cursor-pointer">Note 2</li>
          <li className="p-4 bg-white border border-gray-300 cursor-pointer">Note 3</li>
        </ul>
      </div>
      <div id="note-editor" className="flex-grow p-4">
        <textarea className="w-full h-full p-4 border border-gray-300 resize-none" placeholder="Write your note here..."></textarea>
      </div>
    </main>
  );
}
