import {
  useGetApiSidebarEntriesEntryId,
  useGetApiSidebarEntriesEntryIdSubEntries,
  usePostApiSidebarEntries,
  usePutApiSidebarEntriesEntryId,
} from '@/api'
import { usePrivateEnties } from '@/hooks/workspace'
import clsx from 'clsx'
import dayjs from 'dayjs'
import { useCallback, useState } from 'react'

export const TestApi = () => {
  const privateId = usePrivateEnties()
  const { data: sidebarData } = useGetApiSidebarEntriesEntryId(privateId ?? '')
  const { data: subData, mutate: refreshData } =
    useGetApiSidebarEntriesEntryIdSubEntries(sidebarData?.data.entry?.id ?? '')
  const { trigger: create, isMutating } = usePostApiSidebarEntries()
  const [id2Update, setId2Update] = useState<string>('')
  const { trigger: update } = usePutApiSidebarEntriesEntryId(id2Update ?? '')

  const createEntry = useCallback(() => {
    create({
      parent: sidebarData?.data.entry?.id,
      type: 'note',
      title: 'new sub ' + dayjs().format('YYYY-MM-DD HH:mm:ss'),
    }).finally(() => {
      refreshData()
    })
  }, [create, refreshData, sidebarData?.data.entry?.id])

  const updateName = useCallback(() => {
    if (!id2Update) return
    update({
      title: 'updated ' + dayjs().format('YYYY-MM-DD HH:mm:ss'),
    }).finally(() => {
      setId2Update('')
      refreshData()
    })
  }, [id2Update, refreshData, update])

  return (
    <div className="relative">
      <div
        className={clsx(
          'absolute h-full w-full flex items-center justify-center bg-slate-400 bg-opacity-85',
          { hidden: !isMutating },
        )}
      >
        <span className="loading loading-dots loading-xl"></span>
      </div>

      <code className="m-5">{JSON.stringify(sidebarData?.data)}</code>
      <code className="m-5">{JSON.stringify(subData?.data)}</code>

      <div className="m-3">
        <h1>{sidebarData?.data.entry?.title}</h1>
        {subData?.data.sub_entries?.map((sub) => (
          <div
            key={sub.id}
            className="ml-2"
            onClick={() => {
              setId2Update(sub.id ?? '')
            }}
          >
            <h2>
              {sub.title} /id:{sub.id}
            </h2>
          </div>
        ))}
      </div>

      <button className="btn" onClick={createEntry} disabled={isMutating}>
        add a sub
      </button>
      <button
        className="btn"
        onClick={() => {
          updateName()
        }}
      >
        update
      </button>
      <input
        type="text"
        placeholder="Type here"
        className="input input-bordered w-full max-w-xs"
        value={id2Update}
        onChange={(e) => {
          setId2Update(e.target.value)
        }}
      />
    </div>
  )
}
