<script lang="ts">
  import BillItem from "../../components/BillItem.svelte";
  import type { Bill, BillItemT } from "$lib/bill";
  let bill: Bill;
</script>

<div>
  <BillItem />
</div>
<button
  class="cursor-pointer fixed bottom-12 left-4 bg-orange-300 size-10 rounded-full flex justify-center items-center"
>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    fill="none"
    viewBox="0 0 24 24"
    stroke-width="1.5"
    stroke="currentColor"
    class="size-6"
  >
    <path
      stroke-linecap="round"
      stroke-linejoin="round"
      d="M12 4.5v15m7.5-7.5h-15"
    />
  </svg>
</button>

<button
  class="cursor-pointer fixed bottom-12 right-4 bg-orange-300 size-10 rounded-full flex justify-center items-center"
>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    fill="none"
    viewBox="0 0 24 24"
    stroke-width="1.5"
    stroke="currentColor"
    class="size-6"
  >
    <path
      stroke-linecap="round"
      stroke-linejoin="round"
      d="m4.5 12.75 6 6 9-13.5"
    />
  </svg>
</button>
