import 'package:flutter/material.dart';

class CounterWidget extends StatelessWidget {
  final int count;
  final VoidCallback onPressed;

  const CounterWidget({super.key, required this.count, required this.onPressed});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        IconButton(
          icon: const Icon(Icons.favorite),
          color: Colors.pink,
          onPressed: onPressed,
        ),
        Text('$count', style: const TextStyle(fontSize: 20)),
      ],
    );
  }
}
