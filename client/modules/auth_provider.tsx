import { useState, createContext, useEffect } from 'react'
import { useRouter } from 'next/router'

export type UserInfo = {
    username: string
    id: string
}

export const AuthContext = createContext<{
    authenticated: boolean
    setAuthenticated: (auth: boolean) => void
    user: UserInfo
    setUser: (user: UserInfo) => void
}>({
    authenticated: false,
    setAuthenticated: () => {},
    user: { username: '', id: '' },
    setUser: () => {},
})

const AuthContextProvider = ({ children }: { children: React.ReactNode }) => {
    const [authenticated, setAuthenticated] = useState(false)
    const [user, setUser] = useState<UserInfo>({ username: '', id: '' })

    const router = useRouter()

    useEffect(() => {
        //從localStorage查看有沒有這個人，若沒有他的資料，則返回login頁面

        const userInfo = localStorage.getItem('user_info')

        if (!userInfo) {
            if (window.location.pathname != '/signup') {
                router.push('/login')
                return
            }
        } else {
            //若在signup頁面，且有此人則把用戶設為當前的那個人的資料
            const user: UserInfo = JSON.parse(userInfo)
            if (user) {
                setUser({
                    username: user.username,
                    id: user.id,
                })
            }
            setAuthenticated(true)
        }
    }, [authenticated])

    return (
        <AuthContext.Provider
            value={{
                authenticated: authenticated,
                setAuthenticated: setAuthenticated,
                user: user,
                setUser: setUser,
            }}
        >
            {children}
        </AuthContext.Provider>
    )
}

export default AuthContextProvider