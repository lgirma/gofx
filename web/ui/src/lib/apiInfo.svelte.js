let env = $state('');
let version = $state('');

export function getApiInfo() {

    return {
        get version() {return version},
        set version(val) {version = val},

        get env() {return env},
        set env(val) {env = val},
    }
}

