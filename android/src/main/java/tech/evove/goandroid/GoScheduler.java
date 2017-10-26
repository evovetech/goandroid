package tech.evove.goandroid;

import io.reactivex.Scheduler;

public class GoScheduler extends Scheduler {
    @Override
    public Worker createWorker() {
        // TODO: pooling
        return new GoWorker(new core.GoWorker());
    }
}
