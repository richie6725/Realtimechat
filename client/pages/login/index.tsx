import {useState,useContext,useEffect} from 'react'
import {API_URL} from "@/constants";
import {useRouter}from 'next/router'
import {AuthContext,UserInfo}from '@/modules/auth_provider'

const index = () => {
    //將輸入值
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const{authenticated}=useContext(AuthContext)
    const router=useRouter()

    // 若使用者通過身分認證，回傳到首頁
    useEffect(()=>{
        if (authenticated){
            router.push('/')
            return
        }
    },[authenticated])

    //設定傳輸email,password數值給後端
    const submitHandler=async(e: React.SyntheticEvent)=>{
        e.preventDefault();
        try{
            const res=await fetch(`${API_URL}/login`,{
                method: "POST",
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({email, password})
            })

            const data=await res.json()
            if (res.ok){
                //從UserInfo設定:modules/auth_provider
                //從localStorage來確認是否有user_info
                const user:UserInfo={
                    username:data.username,
                    id:data.id,
                }
                localStorage.setItem('user_info',JSON.stringify(user))
                return router.push('/')
            }

        } catch (err){
            console.log(err)
        }
    }

    return (
        <div className='flex items-center justify-center min-w-full min-h-screen'>
            <form className="flex flex-col md:w-1/5">
                <div className='text-3xl fromt-bold text-center'>
                    <span className='text-blue'>Welcome to chat!</span>
                </div>

                {/*每當使用者在輸入框中輸入內容，會觸發 onChange 事件執行 setEmail，更新狀態值*/}
                <input type='email' placeholder='email'
                       className='p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue'
                       value={email}
                       onChange={(e) => setEmail(e.target.value)}
                />

                {/*傳輸email欄位數值給setEmail*/}
                <input type='password' placeholder='password'
                       className='p-3 mt-4 rounded-md border-2 border-grey focus:outline-none focus:border-blue'
                       value={password}
                       onChange={(e) => setPassword(e.target.value)}
                />
                {/*設定按鈕的部分*/}
                <button className='p-3 mt-6 rounded-md bg-blue font-bold text-white' type='submit' onClick={submitHandler}>
                    login
                </button>
            </form>
        </div>
    )
}
export default index
