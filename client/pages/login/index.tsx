import { useState, useContext, useEffect } from 'react';
import { useRouter } from 'next/router';
import { API_URL } from '@/constants';
import { AuthContext, UserInfo } from '@/modules/auth_provider';  // 確保有引入 AuthContext

const LoginPage = () => {
    //將輸入值
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const { authenticated } = useContext(AuthContext);
    const router = useRouter();

    // 若使用者通過身分認證，回傳到首頁
    useEffect(() => {
        if (authenticated) {
            router.push('/');
            return;
        }
    }, [authenticated,router]);  // ✅ 修正依賴項，確保只在 `authenticated` 變化時觸發

    //設定傳輸email,password數值給後端
    const submitHandler = async (e: React.SyntheticEvent) => {
        e.preventDefault();
        try {
            const res = await fetch(`${API_URL}/login`, {  // ✅ 修正 API_URL，確保請求正確
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, password }),
            });

            const data = await res.json();
            if (res.ok) {
                //從UserInfo設定:modules/auth_provider
                //從localStorage來確認是否有user_info
                const user: UserInfo = {
                    username: data.username,
                    id: data.id,
                };
                localStorage.setItem('user_info', JSON.stringify(user)); // ✅ 修正物件格式
                router.push('/');
            }
        } catch (err) {
            console.log(err);
        }
    };

    return (
        <div className='flex items-center justify-center min-w-full min-h-screen'>
            <form className="flex flex-col md:w-1/5">
                <div className='text-3xl font-bold text-center'>
                    <span className='text-blue'>Welcome to Chat!</span>
                </div>
                {/*每當使用者在輸入框中輸入內容，會觸發 onChange 事件執行 setEmail，更新狀態值*/}
                <input
                    type='email'
                    placeholder='Email'
                    className='p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue'
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />
                {/*傳輸email欄位數值給setEmail*/}
                <input
                    type='password'
                    placeholder='Password'
                    className='p-3 mt-4 rounded-md border-2 border-grey focus:outline-none focus:border-blue'
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                {/*設定按鈕的部分*/}
                <button
                    className='p-3 mt-6 rounded-md bg-blue font-bold text-white'
                    type='submit'
                    onClick={submitHandler}
                >
                    Login
                </button>
            </form>
        </div>
    );
};

export default LoginPage;
