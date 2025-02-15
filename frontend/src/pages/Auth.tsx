import { createEffect, createSignal, Show } from "solid-js"

export default function Auth() {
  function authenticate(){
    alert("Hey")
  }
  const passwordHash = "$2a$12$Rz9JbWdElq6nM.4Nn7a0yOJxWp8naTQsD5NU7F84PWOQ22xbrIUHa"
  let [pwd, setPwd] = createSignal("")
  const [loading, setLoading] = createSignal(true)
  createEffect(() => {
    console.log("hey")
    let isAuthed = localStorage.getItem("isAuthed")
    console.log(isAuthed)
    if (isAuthed) {
      alert("we authed")
    } else {
      setLoading(() => false)
    }
  })

  return (
    <>
      <Show when={loading()}>
        <h1>Loading</h1>
      </Show>
      <Show when={!loading()}>
        <main class=" w-full h-screen flex justify-center items-center">
          <div class="flex flex-col items-center gap-2">
            <input value={pwd()} onInput={(e) => setPwd(e.target.value)} class="border p-2 rounded-md text-center" placeholder="Password" />
            <button onClick={authenticate} class="w-full bg-black text-white p-2 rounded-md">Authenticate</button>
          </div>
        </main>
      </Show>
    </>
  )
}
