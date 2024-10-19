import { useCallback, useState, useRef, useEffect } from 'react'
import { Link } from 'react-router-dom'
import Sortable from 'sortablejs'
import dayjs from 'dayjs'
import clsx from 'clsx'
import {
  useDeleteApiSidebarEntriesEntryId,
  useGetApiSidebarEntriesEntryId,
  useGetApiSidebarEntriesEntryIdSubEntries,
  usePostApiSidebarEntries,
  usePutApiSidebarEntriesEntryId,
} from '@/api'
import { usePrivateEnties } from '@/hooks/workspace'

interface HandlersEntry {
  id?: string
  title?: string
}

export const Sidebar = () => {
  const privateId = usePrivateEnties()
  const { data: sidebarData } = useGetApiSidebarEntriesEntryId(privateId ?? '')
  const { data: subData, mutate: refreshData } =
    useGetApiSidebarEntriesEntryIdSubEntries(sidebarData?.data.entry?.id ?? '')
  const { trigger: create, isMutating: isCreating } = usePostApiSidebarEntries()
  const [selectedId, setSelectedId] = useState<string>('')
  const { trigger: update, isMutating: isUpdating } =
    usePutApiSidebarEntriesEntryId(selectedId ?? '')
  const { trigger: deleteEntity, isMutating: isDeleting } =
    useDeleteApiSidebarEntriesEntryId(selectedId ?? '')
  const [openSubmenuId, setOpenSubmenuId] = useState<string | null>(null)
  const sortableRef = useRef<HTMLUListElement>(null)
  const submenuRef = useRef<HTMLDivElement>(null)

  const isLoading = isCreating || isUpdating || isDeleting

  useEffect(() => {
    if (sortableRef.current) {
      Sortable.create(sortableRef.current, {
        animation: 150,
        onEnd: (evt) => {
          const { oldIndex, newIndex } = evt
          if (oldIndex !== newIndex) {
            console.log(`Moved item from ${oldIndex} to ${newIndex}`)
            // Here you would typically update the order in your backend
          }
        },
      })
    }
  }, [])

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        submenuRef.current &&
        !submenuRef.current.contains(event.target as Node)
      ) {
        setOpenSubmenuId(null)
      }
    }

    document.addEventListener('mousedown', handleClickOutside)
    return () => {
      document.removeEventListener('mousedown', handleClickOutside)
    }
  }, [])

  const createEntry = useCallback(() => {
    create({
      parent: sidebarData?.data.entry?.id,
      type: 'note',
      title: 'New Entry ' + dayjs().format('YYYY-MM-DD HH:mm:ss'),
    }).finally(() => {
      refreshData()
      setOpenSubmenuId(null)
    })
  }, [create, refreshData, sidebarData?.data.entry?.id])

  const updateEntry = useCallback(() => {
    if (!selectedId) return
    update({
      title: 'Updated ' + dayjs().format('YYYY-MM-DD HH:mm:ss'),
    }).finally(() => {
      refreshData()
      setOpenSubmenuId(null)
    })
  }, [selectedId, update, refreshData])

  const deleteEntry = useCallback(() => {
    if (!selectedId) return
    deleteEntity().finally(() => {
      refreshData()
      setSelectedId('')
      setOpenSubmenuId(null)
    })
  }, [deleteEntity, selectedId, refreshData])

  const toggleSubmenu = (id: string) => {
    setOpenSubmenuId(openSubmenuId === id ? null : id)
    setSelectedId(id)
  }

  return (
    <div className="drawer-open relative">
      <input id="my-drawer-2" type="checkbox" className="drawer-toggle" />
      <div
        className={clsx(
          'absolute inset-0 z-50 flex items-center justify-center bg-slate-400 bg-opacity-85',
          { hidden: !isLoading },
        )}
      >
        <span className="loading loading-dots loading-xl"></span>
      </div>
      <div className="drawer-side">
        <label
          htmlFor="my-drawer-2"
          aria-label="close sidebar"
          className="drawer-overlay"
        ></label>
        <div className="bg-base-200 text-base-content min-h-full w-80 p-4 flex flex-col">
          <ul ref={sortableRef} className="space-y-2">
            {subData?.data.sub_entries?.map((entry: HandlersEntry) => (
              <li key={entry.id} className="relative">
                <div className="flex items-center justify-between p-2 hover:bg-base-300 rounded">
                  <Link
                    to={`/article/${entry.id}`}
                    className={`flex-grow ${
                      selectedId === entry.id ? 'font-bold' : ''
                    }`}
                    onClick={() => setSelectedId(entry.id ?? '')}
                  >
                    {entry.title}
                  </Link>
                  <button
                    className="text-gray-500 hover:text-gray-700"
                    onClick={() => entry.id && toggleSubmenu(entry.id)}
                  >
                    ...
                  </button>
                </div>
                {openSubmenuId === entry.id && (
                  <div
                    ref={submenuRef}
                    className="absolute right-0 mt-2 w-48 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 z-10"
                  >
                    <div
                      className="py-1"
                      role="menu"
                      aria-orientation="vertical"
                      aria-labelledby="options-menu"
                    >
                      <button
                        onClick={createEntry}
                        className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left"
                        role="menuitem"
                      >
                        New page
                      </button>
                      <button
                        onClick={updateEntry}
                        className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left"
                        role="menuitem"
                      >
                        Update
                      </button>
                      <button
                        onClick={deleteEntry}
                        className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left"
                        role="menuitem"
                      >
                        Delete
                      </button>
                    </div>
                  </div>
                )}
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  )
}
