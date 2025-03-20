<script lang="ts">
  import Product from "../components/Product.svelte";
  import { products } from "$lib/database"
</script>

<div class=" w-full h-auto flex-col items-center gap-8">
  {#each products as product }
    <Product data={product}/>
  {/each}
</div>
