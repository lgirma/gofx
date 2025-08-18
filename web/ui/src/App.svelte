<script>
  import "./app.css";
  import {ModeWatcher, toggleMode} from "mode-watcher";
  import MainPage from "./MainPage.svelte";
  import {Pages} from "./services/pages.js";
  import {Button} from "$lib/components/ui/button/index.js";
  import {Separator} from "$lib/components/ui/separator/index.js";
  import {Progress} from "$lib/components/ui/progress/index.js";
  import * as Card from "$lib/components/ui/card/index.js";
  import Sun from "@lucide/svelte/icons/sun";
  import Moon from "@lucide/svelte/icons/moon";
  import ActivationPage from "$lib/components/ui/ActivationPage.svelte";
  import {onMount} from "svelte";
  import {isLicensed} from "$lib/licensing.js";
  import {Badge} from "$lib/components/ui/badge/index.js";
  import {getApiInfo} from "$lib/apiInfo.svelte.js";

  let installed = $state(null);
  let info = getApiInfo();

  onMount(async () => {
    const result = await isLicensed();
    installed = result.data;
  });

</script>

<main class="flex align-middle items-center justify-center" style="height: 100vh">
  <div class="flex flex-col" style="max-width: 1024px;">

    <Card.Root class="w-[400px]">
      <Card.Header>
        <Card.Title>
          <div class="flex align-middle items-center gap-3">
            <img src="favicon.png" alt="logo" height="32" width="32" class="inline" />
            <span class="flex-grow flex">
              Import EJ Log
              {#if info.version}
                <Badge variant="ghost" class="ml-2">{info.version}</Badge>
              {/if}
            </span>
            <Button onclick={toggleMode} variant="outline" size="icon">
              <Moon
                  class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
              />
              <Sun
                  class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
              />
            </Button>
          </div>
        </Card.Title>
        <Card.Description>Import/export from electronic journal log files.</Card.Description>
      </Card.Header>
      <Card.Content>
        {#if installed == null}
          <div>Loading ...</div>
        {:else if installed === false}
          <ActivationPage />
        {:else}
          <MainPage />
        {/if}
      </Card.Content>
    </Card.Root>

    <ModeWatcher />

  </div>
</main>
