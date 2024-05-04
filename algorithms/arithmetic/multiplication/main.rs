fn main() {
    let mut x = 7;
    let mut y = 10;
    let mut z = 0;

    while y > 0 {
        if y % 2 != 0 {
            y = y - 1;
            z = z + x;
        } else {
            y = y / 2;
            x = x + x;
        }
    }

    println!("Ergebnis: {z}")
}
