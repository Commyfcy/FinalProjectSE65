import NextAuth from "next-auth"
import CredentialsProvider from "next-auth/providers/credentials"
export const authOptions = {
providers: [
CredentialsProvider({
name: 'Credentials',
credentials: {
username: { label: "Username", type: "text" },
password: { label: "Password", type: "password" }
},
async authorize(credentials, req) {
const res = await fetch("https://127.0.0.1:8000/login", {
method: 'POST',
body: JSON.stringify(credentials),
headers: { "Content-Type": "application/json" }
})
const data = await res.json()
if(data.status == 'ok'){
return data.user
}
return null
}
})
]
}
export default NextAuth(authOptions)