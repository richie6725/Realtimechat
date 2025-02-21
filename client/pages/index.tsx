import { useState, useEffect, useContext, useCallback } from 'react'
import { API_URL } from '../constants'
import { v4 as uuidv4 } from 'uuid'
import { WEBSOCKET_URL } from '../constants'
import { AuthContext } from '../modules/auth_provider'
import { WebsocketContext } from '../modules/websocket_provider'
import { useRouter } from 'next/router'

const Home = () => {
    const [rooms, setRooms] = useState<{ id: string; name: string }[]>([])
    const [roomName, setRoomName] = useState('')
    const { user } = useContext(AuthContext)
    const { setConn } = useContext(WebsocketContext)
    const router = useRouter()

    // ✅ 用 useCallback 來包裝，確保 useEffect 依賴不會變化
    const getRooms = useCallback(async () => {
        try {
            console.log('Fetching rooms...') // ✅ 確保 API 有被呼叫
            const res = await fetch(`${API_URL}/ws/getRooms`, {
                method: 'GET',
            })
            const data = await res.json()
            if (res.ok) {
                console.log('Rooms fetched:', data) // ✅ 確保有收到資料
                setRooms(data)
            } else {
                console.error('Failed to fetch rooms:', res.status)
            }
        } catch (err) {
            console.error('Error fetching rooms:', err)
        }
    }, [])

    // ✅ 修正 useEffect 依賴，確保只在 `conn` 變化時重新獲取
    useEffect(() => {
        getRooms()
    }, [getRooms])

    const submitHandler = async (e: React.SyntheticEvent) => {
        e.preventDefault()

        if (!roomName.trim()) {
            console.warn('Room name cannot be empty') // ✅ 確保 roomName 不是空值
            return
        }

        const newRoom = { id: uuidv4(), name: roomName.trim() }
        setRoomName('') // ✅ 清空輸入框，但不影響 `fetch` 的 roomName

        try {
            console.log('Creating room:', newRoom) // ✅ 確保 API 有被呼叫
            const res = await fetch(`${API_URL}/ws/createRoom`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
                body: JSON.stringify(newRoom),
            })

            if (res.ok) {
                console.log('Room created successfully') // ✅ 確保 API 成功
                getRooms() // ✅ 重新獲取房間
            } else {
                console.error('Failed to create room:', res.status)
            }
        } catch (err) {
            console.error('Error creating room:', err)
        }
    }

    const joinRoom = (roomId: string) => {
        if (!user) {
            console.error('User not authenticated') // ✅ 確保 user 存在
            return
        }

        const ws = new WebSocket(
            `${WEBSOCKET_URL}/ws/joinRoom/${roomId}?userId=${user.id}&username=${user.username}`
        )

        ws.onopen = () => {
            console.log('WebSocket connected:', ws) // ✅ 確保 WebSocket 連線成功
            setConn(ws)
            router.push('/app')
        }

        ws.onerror = (err) => {
            console.error('WebSocket error:', err) // ✅ 確保錯誤能被偵測
        }
    }

    return (
        <>
            <div className='my-8 px-4 md:mx-32 w-full h-full'>
                <div className='flex justify-center mt-3 p-5'>
                    <input
                        type='text'
                        className='border border-grey p-2 rounded-md focus:outline-none focus:border-blue'
                        placeholder='room name'
                        value={roomName}
                        onChange={(e) => setRoomName(e.target.value)}
                    />
                    <button
                        className='bg-blue border text-white rounded-md p-2 md:ml-4'
                        onClick={submitHandler}
                    >
                        create room
                    </button>
                </div>
                <div className='mt-6'>
                    <div className='font-bold'>Available Rooms</div>
                    <div className='grid grid-cols-1 md:grid-cols-5 gap-4 mt-6'>
                        {rooms.map((room, index) => (
                            <div
                                key={index}
                                className='border border-blue p-4 flex items-center rounded-md w-full'
                            >
                                <div className='w-full'>
                                    <div className='text-sm'>room</div>
                                    <div className='text-blue font-bold text-lg'>{room.name}</div>
                                </div>
                                <div className=''>
                                    <button
                                        className='px-4 text-white bg-blue rounded-md'
                                        onClick={() => joinRoom(room.id)}
                                    >
                                        join
                                    </button>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            </div>
        </>
    )
}

export default Home
