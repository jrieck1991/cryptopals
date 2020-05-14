fn main() {

    let hex_str = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736";
    let byte_string = ascii_to_hex_bytes(hex_str.to_string()); 
    println!("{:?}", String::from_utf8(byte_string));
}


fn ascii_to_hex_bytes(hex_str: String) -> Vec<u8> {

    // init result vector
    let mut result: Vec<u8> = Vec::new();

    // get length
    let len = hex_str.chars().count();

    // some validation that it's a hex string
    if len % 2 != 0 {
        return result
    };

    // convert to byte vector
    let hex_bytes: Vec<u8> = hex_str.into_bytes();

    // iterate over bytes
    for (i, _b) in hex_bytes.iter().enumerate().step_by(2) {

        let mut b1: u8 = hex_bytes[i]; 

        // if byte is > hex 'A', subtract by hex7 to get 0th
        if b1 >= 52 {
            b1 = b1 - 49
        };

        // shift last 4 bits to high bits
        let high_nibble: u8 = b1 << 4;

        let mut b2: u8 = hex_bytes[i+1];

        if b2 >= 52 {
            b2 = b2 - 49
        };

        // grab last 4 bits with mask
        let low_nibble: u8 = b2 & 0x0F;

        // combine
        let byte = high_nibble | low_nibble;

        result.push(byte)
    };

    return result
}