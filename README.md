# Rust Depot Action

A GitHub action that will query [steamcmd.net](https://www.steamcmd.net/) for the latest public build id for Rust Dedicated Server, along with the Linux & Common manifest IDs.

## Inputs

|Name|Default|Description|
|---|---|---|
|app_id|258550|The App ID to fetch (Currently only supports Rust Dedicated)|

## Outputs

|Name|Type|Description|
|---|---|---|
|build_id|Number|The Latest Build ID for Rust|
|build_updated_time|Unix Timestamp|When the latest build was updated at|
|common_manifest_id|Number|The Manifest for the 'rust dedicated - linux64' Depot|
|linux_manifest_id|Number|The Manifest for the 'rust dedicated - common' Depot|


## Example
```yaml
      - name: Rust App
        uses: whype-net/rust-depot-action@v1
        id: rust
        with:
          app_id: 258550

      - name: Get output build ID
        run: echo "Latest build id is ${{ steps.rust.outputs.build_id }}"
```
