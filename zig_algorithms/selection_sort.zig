const std = @import("std");
const io_utils = @import("io_utils.zig");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const parsed = try io_utils.readInput(allocator);
    defer parsed.deinit();
    
    var arr = try allocator.alloc(i64, parsed.value.array.len);
    defer allocator.free(arr);
    @memcpy(arr, parsed.value.array);

    var step_num: usize = 0;
    var sorted_indices = std.ArrayList(usize).init(allocator);
    defer sorted_indices.deinit();

    try io_utils.printStep(arr, &.{}, &.{}, sorted_indices.items, step_num, "Starting Selection Sort");
    step_num += 1;

    var i: usize = 0;
    while (i < arr.len - 1) : (i += 1) {
        var min_index = i;
        
        var msg_buf: [100]u8 = undefined;
        const start_msg = try std.fmt.bufPrint(&msg_buf, "Current minimum is {d} at {d}", .{arr[min_index], min_index});
        
        // Visualize starting scan
        try io_utils.printStep(arr, &.{i}, &.{}, sorted_indices.items, step_num, start_msg);
        step_num += 1;

        var j: usize = i + 1;
        while (j < arr.len) : (j += 1) {
            const compare_indices = [_]usize{min_index, j};
            const compare_msg = try std.fmt.bufPrint(&msg_buf, "Comparing minimum ({d}) with {d}", .{arr[min_index], arr[j]});
            try io_utils.printStep(arr, &compare_indices, &.{}, sorted_indices.items, step_num, compare_msg);
            step_num += 1;

            if (arr[j] < arr[min_index]) {
                min_index = j;
                const new_min_msg = try std.fmt.bufPrint(&msg_buf, "Found new minimum: {d}", .{arr[min_index]});
                try io_utils.printStep(arr, &.{min_index}, &.{}, sorted_indices.items, step_num, new_min_msg);
                step_num += 1;
            }
        }

        if (min_index != i) {
            const temp = arr[i];
            arr[i] = arr[min_index];
            arr[min_index] = temp;
            
            const swap_indices = [_]usize{i, min_index};
            const swap_msg = try std.fmt.bufPrint(&msg_buf, "Swapping {d} and {d}", .{arr[i], arr[min_index]});
            try io_utils.printStep(arr, &.{}, &swap_indices, sorted_indices.items, step_num, swap_msg);
            step_num += 1;
        }

        try sorted_indices.append(i);
    }
    
    // Last element is sorted
    try sorted_indices.append(arr.len - 1);
    try io_utils.printStep(arr, &.{}, &.{}, sorted_indices.items, step_num, "Sorting complete!");
}
