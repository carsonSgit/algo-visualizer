const std = @import("std");

var unsorted_array: [5]u8 = undefined;

pub fn populate_array() void {
    var prng = std.Random.DefaultPrng.init(42);
    var random = prng.random();

    for (0..unsorted_array.len) |i| {
        unsorted_array[i] = random.int(u8) % 10;
    }
}

// learnt: zig custom loop syntax
// learnt: zig is very type strict, typecasting is very useful for working with arrays
pub fn insertion_sort() void {
    var i: u8 = 1;

    while (i < unsorted_array.len) : (i += 1) {
        const key = unsorted_array[@as(usize, @intCast(i))];
        var j: i8 = @as(i8, @intCast(i)) - 1;

        while (j >= 0 and unsorted_array[@as(usize, @intCast(j))] > key) : (j -= 1) {
            unsorted_array[@as(usize, @intCast(j + 1))] = unsorted_array[@as(usize, @intCast(j))];
        }

        unsorted_array[@as(usize, @intCast(j + 1))] = key;
    }
}

pub fn main() void {
    populate_array();
    std.debug.print("Unsorted array: {d}\n", .{unsorted_array});
    insertion_sort();
    std.debug.print("Sorted array: {d}\n", .{unsorted_array});
}
