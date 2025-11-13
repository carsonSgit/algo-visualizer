const std = @import("std");

var unsorted_array: [5]u8 = undefined;

pub fn populate_array() void {
    var prng = std.Random.DefaultPrng.init(42);
    var random = prng.random();
    for (0..unsorted_array.len) |i| {
        unsorted_array[i] = random.int(u8) % 10;
    }
}

pub fn bubble_sort() void {
    var temp: u8 = 0;
    var swapped = false;

    // repeat until all elements were sorted
    for (0..unsorted_array.len - 1) |i| {
        swapped = false;
        // do a nested pass through (swapping one element at a time)
        for (0..unsorted_array.len - i - 1) |j| {
            // if current element is greater than next, swap
            if (unsorted_array[j] > unsorted_array[j + 1]) {
                temp = unsorted_array[j];
                unsorted_array[j] = unsorted_array[j + 1];
                unsorted_array[j + 1] = temp;
                swapped = true;
            }
        }
        if (!swapped) break;
    }
}

pub fn main() void {
    populate_array();
    std.debug.print("{d}\n", .{unsorted_array});
    bubble_sort();
    std.debug.print("{d}\n", .{unsorted_array});
}
