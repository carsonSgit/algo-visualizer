const std = @import("std");

var unsorted_array: [5]u8 = undefined;

pub fn populate_array() void {
    var prng = std.Random.DefaultPrng.init(42);
    var random = prng.random();
    for (0..unsorted_array.len) |i| {
        unsorted_array[i] = random.int(u8) % 10;
    }
}

pub fn selection_sort() void {
    // temp variables for swapping
    var temp: u8 = 0;
    var min_index: u8 = 0;

    for (0..unsorted_array.len - 1) |i| {
        // min_index is the smallest index in the unsorted array, default to i as selection sort searches for the smallest element within the unsorted array
        min_index = @as(u8, @intCast(i));
        // search for the smallest element in the unsorted array, starting from i + 1 all the way till the end of the array
        for (i + 1..unsorted_array.len) |j| {
            if (unsorted_array[min_index] > unsorted_array[j]) {
                min_index = @as(u8, @intCast(j));
            }
        }
        temp = unsorted_array[i];
        unsorted_array[i] = unsorted_array[min_index];
        unsorted_array[min_index] = temp;
    }
}

pub fn main() void {
    populate_array();
    std.debug.print("{d}\n", .{unsorted_array});
    selection_sort();
    std.debug.print("{d}\n", .{unsorted_array});
}
