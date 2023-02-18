import { createSignal } from "solid-js"
import './sliders.css'

export const Sliders = () => {
    const [value, setValue] = createSignal("50")

    return (
        <div class="sliders">
  <input type="range" min="1" max="100" value={value()} class="slider" id="myRange" onChange={(e) => setValue(e.currentTarget.value)} />
</div>
    )
}