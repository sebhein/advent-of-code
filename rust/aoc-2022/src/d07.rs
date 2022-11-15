pub fn solve_07(_input: &String) {
    let sol_1 = solve_p1(_input);
    let sol_2 = solve_p2(_input);
    println!("day 07, part 1: {} | part 2: {}", sol_1, sol_2);
}


fn solve_p1(_input: &String) -> u32 {
    let mut total_size = 0;
    let mut tree: Vec<u32> = Vec::new();
    let mut depth = 0;
    for line in _input.lines() {
        match line {
            _ if &line[0..] == "$ cd .." => {
                match tree.pop() {
                    None => println!("probably shouldn't happen"),
                    Some(size) => {
                        if size < 100000 { total_size += size }
                        depth -= 1;
                        tree[depth] += size;
                    }
                }
            },
            _ if &line[0..4] == "$ cd" => {
                if &line[5..] == "/" { 
                    depth = 0;
                } else {
                    depth += 1;
                }

                tree.push(0);
            },
            _ if &line[0..4] == "$ ls" => continue,
            _ if &line[0..1] == "d" => continue,
            _ => tree[depth] += line.split_whitespace().next().unwrap().parse::<u32>().unwrap(),
        }
        //println!("line: {} | total: {} | tree: {:?}", line, total_size, tree);
    }

    // unwind the last depth traversal
    while let Some(size) = tree.pop() {
        //println!("size is {}, total_size {}", size, total_size);
        if size < 100000 { total_size += size } else { break }
        if depth == 0 { break }
        depth -= 1;
        tree[depth] += size;
    }
    return total_size;
}


fn solve_p2(_input: &String) -> u32 {
    let mut tree: Vec<u32> = Vec::new();
    let mut depth = 0;
    let mut dir_sizes: Vec<u32> = Vec::new();
    for line in _input.lines() {
        match line {
            _ if &line[0..] == "$ cd .." => {
                let size = tree.pop().unwrap();
                depth -= 1;
                tree[depth] += size;
                dir_sizes.push(size);
            },
            _ if &line[0..4] == "$ cd" => {
                if &line[5..] == "/" { 
                    depth = 0;
                } else {
                    depth += 1;
                }

                tree.push(0);
            },
            _ if &line[0..4] == "$ ls" => continue,
            _ if &line[0..1] == "d" => continue,
            _ => tree[depth] += line.split_whitespace().next().unwrap().parse::<u32>().unwrap(),
        }
    }

    // unwind the last depth traversal
    while let Some(size) = tree.pop() {
        dir_sizes.push(size);
        if depth == 0 { break }
        depth -= 1;
        tree[depth] += size;
    }

    let total = dir_sizes.pop().unwrap();
    let diff = 30000000 - (70000000 - total);

    dir_sizes.sort();
    for size in dir_sizes {
        if size >= diff {
            return size;
        }
    }

    return 0;
}

#[cfg(test)]
mod tests {
    use super::*;
    const EXAMPLE_INPUT: &str = 
"$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
";

    #[test]
    fn test_p1() {
        assert_eq!(solve_p1(&EXAMPLE_INPUT.to_string()), 95437)
    }

    #[test]
    fn test_last_directory_is_counted() {
        let input: &str =
"$ cd /
$ ls
dir a
1 b.txt
$ cd a 
$ ls
1 c.txt
";
        assert_eq!(solve_p1(&input.to_string()), 3)
    }

    #[test]
    fn test_p2() {
        assert_eq!(solve_p2(&EXAMPLE_INPUT.to_string()), 24933642)
    }
}
