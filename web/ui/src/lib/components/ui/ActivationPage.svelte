<script>
    import {Label} from "$lib/components/ui/label/index.js";
    import {onMount} from "svelte";
    import {activate, getRequestCode} from "../../licensing.js";
    import {Textarea} from "$lib/components/ui/textarea/index.js";
    import {Button} from "$lib/components/ui/button/index.js";
    import {Input} from "$lib/components/ui/input/index.js";

    let reqCode = $state('');
    let activationCode = $state('');
    let activationAttemptStatus = $state(null);
    let busy = $state(false);
    onMount(async () => {
        busy = true;
        const response = await getRequestCode();
        reqCode = response.data;
        busy = false;
    });

    async function onActivate() {
        busy = true;
        const response = await activate(activationCode);
        activationAttemptStatus = !response.error;
        busy = false;
    }
</script>

<div class="flex flex-col gap-2">
    {#if busy}
        <div class="text-center text-gray-500 font-bold">
            Checking ...
        </div>
    {:else if activationAttemptStatus == null}
        <div class="text-center text-orange-500 font-bold">
            Activate License
        </div>
    {:else if activationAttemptStatus}
        <div class="text-center text-green-500 font-bold">
            Activation succeeded
        </div>
    {:else}
        <div class="text-center text-red-500 font-bold">
            Activation failed
        </div>
    {/if}

    <div>
        <Label>Request Code</Label>
        <Input value={reqCode} disabled class="bg-gray-400 text-sm" />
    </div>
    <div>
        <Label>Activation Code</Label>
        <Textarea bind:value={activationCode} class="text-sm" />
    </div>
    <div class="text-center">
        {#if activationAttemptStatus}
            <a href="/public">Start</a>
        {:else}
            <Button onclick={onActivate}>Activate</Button>
        {/if}
    </div>
</div>
