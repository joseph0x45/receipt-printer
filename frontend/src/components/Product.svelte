<script lang="ts">
  import { type ProductT } from "$lib/database";
  let data = $props() as ProductT;
</script>

<div class="p-4 flex flex-col gap-6 w-full bg-sky-200 mb-4">
  <img class="w-92 h-92" src="https://picsum.photos/200/300" alt="" />
  <div class="w-full flex flex-col gap-4">
    <h2>Name</h2>
    <p>Price</p>
    <p>Quantity in stock</p>
  </div>
</div>
