# Bitcoin HD Wallet Generator


### What Is an HD Wallet?

An HD Wallet, or Hierarchical Deterministic wallet, is a new-age digital wallet that automatically generates a hierarchical tree-like structure of private/public addresses (or keys), thereby addressing the problem of the user having to generate them on their own.


### HD Wallet Explained
A standard cryptocurrency wallet is used to store the cryptocurrency tokens or coins. It has a public address which the user can give to others to receive funds from them, and a private key that the user uses to spend the stored tokens.

This public/private combination mechanism ensures safety of the cryptocurrency tokens but comes with an additional overhead of the user being required to repeatedly generate a random pair of private/public addresses (or keys), and back up each time one configures a new pair of addresses. As the number of transactions increases, this process becomes cumbersome for the user.

HD Wallets, or Hierarchical Deterministic wallets, solve this problem by deriving all the addresses from a single master seed (hence the name hierarchical). All HD wallets use a variant of the standard 12-word master seed key, and each time this seed is extended at the end by a counter value which makes it possible to automatically derive an unlimited number of new addresses.

HD wallets eliminate the need for the user to constantly generate and wait for the secure keys to be generated, so they only need to worry about taking the backup.

This repo is continuing the work of https://github.com/foxnut/go-hdwallet
(actually removing lots of stuff form it and making it work just for bitcoin)


### How to use it

- download the files from the hdwallet package, and use it in your code.

### See main.go file for an example

~~~~
~~~~~~~~~~~ BITCOIN MAINNET ~~~~~~~~~~~
WIF (compressed): L4meG9M7FHvd1rVjwbZwJAvd4fHdsweGzC4M1nurDttfRuf1tohW
Address bc1q4950q3q3cxtlve5yq2rd47ksd2qgr20uv4e5jn
Segwit 38H9WZNBWVcMNHvt8CzvnmvzFxMm5QZvbn
Legacy 1GSkugFrs5SLBqmrnRvfUUaKRjurYxs5ji
------------
WIF (compressed): Kwug6rhcbFbkX6RZ9oKv72sefVe5pXAALiMnmMa15qzYQkWFpLJM
Address bc1qhtg2ffqp0q4a8gfaxlqu0xg0j9vc6hxl3e5nd3
Segwit 3GT9YrSWvQoZjNssPcXQMnxPzJA2E1WR7g
Legacy 1J2ndTdRmkE49hrwxvCH52FdbhqMi5q9io
------------
WIF (compressed): Ky1rH9pX5i89zDYz3d1oS83Gg4KzDQTYSnuN8CLS4n22H9ZxMdDg
Address bc1qc9smljrgltplcppc2ewq4405jduu24zp6hseru
Segwit 3E5TKiipcGX8uTpa7CCkymGdhzNQRA3iK3
Legacy 1JdWWkhSu8cDTTUhv89AS8zZmrSGkwHCT5
------------


~~~~~~~~~~~ BITCOIN TESTNET ~~~~~~~~~~~


WIF (compressed): cV8dj4LxgMctBHy1L1P4fVRggtb3YPjy4ECp8DNMj1Yfgen2YhpC
Address tb1q4950q3q3cxtlve5yq2rd47ksd2qgr20uxnz8fq
Segwit 2MyqMaJJD7x7ha5ZRoLcoQivFUJZvrcGYbj
Legacy mvxiCjLqg6saxxFUVzu3JPneHjWZQ5fAbY
------------
WIF (compressed): cNGfZmhU2KJ1gXtpYD93UMNiHiwVUyFrQkWFsn2WaxeYfVX2SZfj
Address tb1qhtg2ffqp0q4a8gfaxlqu0xg0j9vc6hxlml0qkz
Segwit 2N81McbNYXsJuwAWR4k9GyjwfCeNBxgvNuj
Legacy mxYjvWiQamfJvpLZgVAetwTxThS4fYc5Bx
------------
WIF (compressed): cPNqk4pNWmpR9f2FS2pvoSYLJHdPsrZEWq3qEcnwZtg2XtfKkPoK
Address tb1qc9smljrgltplcppc2ewq4405jduu24zps3t2c0
Segwit 2N5dfPTerDj2V7FT7nKpdbiFtvLaaBA5UZe
Legacy my9ToonRiA3UEZxKdh7YG4Ctdr2ygPU14Z
------------
~~~~


### License: MIT
