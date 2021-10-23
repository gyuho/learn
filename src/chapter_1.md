# Chapter 1

TBD

```rust,edition2018
use data_encoding::HEXUPPER;
use ring::digest::{Context, Digest, SHA256};
use std::io::{BufReader, Read};

use error_chain::error_chain;
error_chain! {
    foreign_links {
        Io(std::io::Error);
        Decode(data_encoding::DecodeError);
    }
}

fn compute_hash<R: Read>(mut reader: R) -> Result<Digest> {
    let mut context = Context::new(&SHA256);
    let mut buffer = [0; 1024];

    loop {
        let count = reader.read(&mut buffer)?;
        if count == 0 {
            break;
        }
        context.update(&buffer[..count]);
    }

    Ok(context.finish())
}

fn main() -> Result<()> {
    let digest1 = compute_hash(BufReader::new("hello".as_bytes()))?;
    let digest2 = compute_hash(BufReader::new("hello.".as_bytes()))?;

    println!("{}", HEXUPPER.encode(digest1.as_ref()));
    // 2CF24DBA5FB0A30E26E83B2AC5B9E29E1B161E5C1FA7425E73043362938B9824
    println!("{}", HEXUPPER.encode(digest2.as_ref()));
    // 1589999B0CA6EF8814283026A9F166D51C70A910671C3D44049755F07F2EB910

    // just one more character but totally different hash value

    Ok(())
}
```
